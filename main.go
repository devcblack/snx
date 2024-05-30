package main

import (
	"fmt"
	"encoding/json"
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"main/api"
	"main/util"
	"main/view"
)

type UI struct {
	App *tview.Application
	Login *view.Login
	Search *view.Search
	SearchComponent *view.SearchComponent
	Assets *view.Assets
}

var client *api.Client
var response []byte
var usernameText, passwordText, baseURLText string
var continuationTokenCounter int
var continuationTokens map[int]string

func main() {
	ui := &UI{
		App: tview.NewApplication(),
		Login: view.NewLoginLayout(),
		Search: view.NewSearchLayout(),
		SearchComponent: view.NewSearchComponentLayout(),
		Assets: view.NewAssetsLayout(),
	}

	ui.HandleLogin()
	
	if err := ui.App.SetRoot(ui.Login.Flex, true).Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

func (ui *UI) HandleLogin() {

ui.Login.Log.SetText("Welcome").SetTextColor(tcell.ColorWhite)

usernameField := ui.Login.Form.GetFormItemByLabel("Username")
passwordField := ui.Login.Form.GetFormItemByLabel("Password")
baseUrlField := ui.Login.Form.GetFormItemByLabel("Url")

username, uok := usernameField.(*tview.InputField)
if !uok {
	log.Fatal("username not ok")
}
password, pok := passwordField.(*tview.InputField)
if !pok {
	log.Fatal("password not ok")
}
baseURL, bok := baseUrlField.(*tview.InputField)
if !bok {
	log.Fatal("baseURL not ok")
}

SNX_USERNAME := os.Getenv("SNX_USERNAME")
SNX_PASSWORD := os.Getenv("SNX_PASSWORD")
SNX_BASE_URL := os.Getenv("SNX_BASE_URL")

if SNX_USERNAME != "" {
	username.SetText(SNX_USERNAME)
}
if SNX_PASSWORD != "" {
	password.SetText(SNX_PASSWORD)
}
if SNX_BASE_URL != "" {
	baseURL.SetText(SNX_BASE_URL)
}

ui.App.SetRoot(ui.Login.Flex, true)

ui.Login.Form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
			case tcell.KeyEnter:
				usernameText = username.GetText()
				passwordText = password.GetText()
				baseURLText = baseURL.GetText()
				if usernameText == "" || passwordText == "" || baseURLText == "" {
					ui.Login.Log.SetText("Username or Password or URL is empty").SetTextColor(tcell.ColorRed)
				} else {
					ui.HandleSearch("","")
				}
			}
		return event
	})
}

func (ui *UI) HandleSearch(keywordText, continuationToken string) {

ui.App.SetRoot(ui.Search.Flex, true)
				
ui.Search.Form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
			case tcell.KeyESC:
				ui.HandleLogin()
			case tcell.KeyEnter:

				keywordField := ui.Search.Form.GetFormItemByLabel("Keyword")
				keyword, uok := keywordField.(*tview.InputField)
				if !uok {
					log.Fatal("keyword not ok")
				}
				keywordText = keyword.GetText()

				client = api.NewClient(usernameText, passwordText, baseURLText)

				continuationTokenCounter = 0
				continuationTokens = make(map[int]string)

				ui.HandleSearchComponent(keywordText, continuationToken)
		}
	return event
	})
}

func (ui *UI) HandleSearchComponent(keywordText, continuationToken string) {

ui.SearchComponent.Table.Clear()


err := errors.New("RequestError")
if continuationToken == ""{

	response, err = client.MakeRequest("/service/rest/v1/search?q=" + keywordText, "GET", nil)
} else {
	response, err = client.MakeRequest("/service/rest/v1/search?continuationToken=" + continuationToken + "&q=" + keywordText, "GET", nil)
}

if err != nil {
	log.Fatal(err)
}

var data *util.SearchComponent
err = json.Unmarshal([]byte(response), &data)
if err != nil {
	log.Printf("could not unmarshal json: %s\n", err)
	return
}

headerlist := []string {"ID", "Repository", "Format", "Group", "Name", "Version"}
for index, header := range headerlist {
	var cell = tview.NewTableCell(header)
	cell.SetTextColor(tcell.ColorYellowGreen)
	ui.SearchComponent.Table.SetCell(0, index, cell).SetSelectable(false, false)

}

for index, item := range data.Items {
	index += 1 // because header is 0

	ui.SearchComponent.Table.SetCell(index, 0, tview.NewTableCell(item.ID)).SetSelectable(true, false)
	ui.SearchComponent.Table.SetCell(index, 1, tview.NewTableCell(item.Repository)).SetSelectable(true, false)
	ui.SearchComponent.Table.SetCell(index, 2, tview.NewTableCell(item.Format)).SetSelectable(true, false)
	ui.SearchComponent.Table.SetCell(index, 3, tview.NewTableCell(item.Group)).SetSelectable(true, false)
	ui.SearchComponent.Table.SetCell(index, 4, tview.NewTableCell(item.Name)).SetSelectable(true, false)	
	ui.SearchComponent.Table.SetCell(index, 5, tview.NewTableCell(item.Version)).SetSelectable(true, false)
}



ui.App.SetRoot(ui.SearchComponent.Flex, true)

ui.SearchComponent.Table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

	switch event.Key() {
		case tcell.KeyESC:
			ui.HandleSearch(keywordText, continuationToken)
		case tcell.KeyEnter:
			ui.SearchComponent.Table.SetSelectedFunc(func(row int, column int) {
				if row > 0 {
					ui.HandleAssets(keywordText, continuationToken, (row - 1))
				}
			})
		}
	switch event.Rune() {
		case 110://n - next
			continuationTokens[continuationTokenCounter] = continuationToken
			continuationTokenCounter += 1
			continuationTokens[continuationTokenCounter] = data.ContinuationToken
			continuationToken = continuationTokens[continuationTokenCounter]

			ui.HandleSearchComponent(keywordText, continuationToken)
		case 78://N - previous
			continuationTokenCounter -= 1
			continuationToken = continuationTokens[continuationTokenCounter]

			ui.HandleSearchComponent(keywordText, continuationToken)
	}
	return event
	})
}


func (ui *UI) HandleAssets(keywordText, continuationToken string, row int) {

ui.Assets.Table.Clear()

client = api.NewClient(usernameText, passwordText, baseURLText)
err := errors.New("RequestError")
if continuationToken == ""{

	response, err = client.MakeRequest("/service/rest/v1/search?q=" + keywordText, "GET", nil)
} else {
	response, err = client.MakeRequest("/service/rest/v1/search?continuationToken=" + continuationToken + "&q=" + keywordText, "GET", nil)
}

if err != nil {
	log.Fatal(err)
}

var data *util.SearchComponent
err = json.Unmarshal([]byte(response), &data)
if err != nil {
	log.Printf("could not unmarshal json: %s\n", err)
	return
}

headerlist := []string {"DownloadURL", "Path", "ID", "Repository", "Format", "LastModified", "LastDownloaded", "Uploader", "UploaderIP", "FileSize", "BlobCreated" }
for index, header := range headerlist {
	var cell = tview.NewTableCell(header)
	cell.SetTextColor(tcell.ColorYellowGreen)
	ui.Assets.Table.SetCell(0, index, cell).SetSelectable(false, false)
}

for index, asset := range data.Items[row].Assets {
	index += 1 // because header is 0
	ui.Assets.Table.SetCell(index, 0, tview.NewTableCell(asset.DownloadURL)).SetSelectable(true, false)
	ui.Assets.Table.SetCell(index, 1, tview.NewTableCell(asset.Path)).SetSelectable(true, false)
	ui.Assets.Table.SetCell(index, 2, tview.NewTableCell(asset.ID)).SetSelectable(true, false)
	ui.Assets.Table.SetCell(index, 3, tview.NewTableCell(asset.Repository)).SetSelectable(true, false)
	ui.Assets.Table.SetCell(index, 4, tview.NewTableCell(asset.Format)).SetSelectable(true, false)
	ui.Assets.Table.SetCell(index, 5, tview.NewTableCell(asset.LastModified.String())).SetSelectable(true, false)
	ui.Assets.Table.SetCell(index, 6, tview.NewTableCell(asset.LastDownloaded.String())).SetSelectable(true, false)
	ui.Assets.Table.SetCell(index, 7, tview.NewTableCell(asset.Uploader)).SetSelectable(true, false)
	ui.Assets.Table.SetCell(index, 8, tview.NewTableCell(asset.UploaderIP)).SetSelectable(true, false)
	ui.Assets.Table.SetCell(index, 9, tview.NewTableCell(strconv.Itoa(asset.FileSize))).SetSelectable(true, false)
	ui.Assets.Table.SetCell(index, 10, tview.NewTableCell(asset.BlobCreated.String())).SetSelectable(true, false)



}

ui.App.SetRoot(ui.Assets.Flex, true)

ui.Assets.Table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

	switch event.Key() {
		case tcell.KeyESC:

			ui.Assets.Log.SetText("ENTER to download Asset").SetTextColor(tcell.ColorWhite)
			ui.HandleSearchComponent(keywordText, continuationToken)
		case tcell.KeyEnter:
			ui.Assets.Table.SetSelectedFunc(func(row int, column int) {
				format := ui.Assets.Table.GetCell(row, 4).Text
				if format == "maven2" || format == "helm" {

					if row > 0 && column == 0 {
						downloadURL := ui.Assets.Table.GetCell(row, column).Text
						
						path := ui.Assets.Table.GetCell(row, 1).Text
						re := regexp.MustCompile(`^.*\/(.*)$`)
						path = re.ReplaceAllString(path, "$1")
						
						err := client.Download(downloadURL, "GET", path, nil)
						if err != nil {
							log.Fatal(err)
						}
						ui.Assets.Log.SetText("Successful download of " + path).SetTextColor(tcell.ColorGreen)
					}
				} else {
					ui.Assets.Log.SetText("No download available of format " + format).SetTextColor(tcell.ColorRed)
				}
			})
		}
	return event
	})
}
