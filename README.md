# snx
slim terminal search and download tool for sonatype-nexus3 in go

### Login
![login-example](https://github.com/devcblack/snx/assets/94953425/8e57df3c-ade6-4c31-af33-7554759f72a5)

### Search
![search-example](https://github.com/devcblack/snx/assets/94953425/5e0f2411-ef98-4c82-b44b-a8f87e838205)

### Component-View
![component-view-example](https://github.com/devcblack/snx/assets/94953425/d808d8b8-8a0f-46a4-9461-f570d510705a)

### Asset-Download
![asset-download-example](https://github.com/devcblack/snx/assets/94953425/a3cb3ef4-3529-4a5a-be1a-d9af9afd1858)


# Installation
### Clone Repository
```
git clone https://github.com/devcblack/snx.git
cd snx
```
### Enable Locale Module
```
export GO111MODULE=on
```
### Build
```
go build -o $HOME/.local/bin/snx
```
### Create Alias In Your .bashrc
```
alias snx="$HOME/.local/bin/snx"
```
### Start snx
```
snx
```


# ENV
* SNX_USERNAME : Username for Login
* SNX_PASSWORD : Password for Login
* SNX_BASE_URL : Base-URL like https://mycompanynexus.com


# Implemented Formats For Download
* maven2
* helm
