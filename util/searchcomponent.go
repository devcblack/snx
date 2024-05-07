package util

import	"time"

type SearchComponent struct {
	Items []struct {
		ID         string `json:"id"`
		Repository string `json:"repository"`
		Format     string `json:"format"`
		Group      string `json:"group"`
		Name       string `json:"name"`
		Version    string `json:"version"`
		Assets     []struct {
			DownloadURL string `json:"downloadUrl"`
			Path        string `json:"path"`
			ID          string `json:"id"`
			Repository  string `json:"repository"`
			Format      string `json:"format"`
			Checksum    struct {
				Sha1   string `json:"sha1"`
				Sha256 string `json:"sha256"`
				Sha512 string `json:"sha512"`
				Md5    string `json:"md5"`
			} `json:"checksum"`
			ContentType    string      `json:"contentType"`
			LastModified   time.Time   `json:"lastModified"`
			LastDownloaded time.Time   `json:"lastDownloaded"`
			Uploader       string      `json:"uploader"`
			UploaderIP     string      `json:"uploaderIp"`
			FileSize       int         `json:"fileSize"`
			BlobCreated    time.Time   `json:"blobCreated"`
			Maven2         struct {
				Extension  string `json:"extension"`
				GroupID    string `json:"groupId"`
				ArtifactID string `json:"artifactId"`
				Version    string `json:"version"`
			} `json:"maven2"`
		} `json:"assets"`
	} `json:"items"`
	ContinuationToken string `json:"continuationToken"`
}
