# snx
slim terminal search and download tool for sonatype-nexus3 in go

### Login
![login-example](https://github.com/devcblack/snx/assets/94953425/f817efdf-f888-40f9-b3a5-5542e5505a7a)
### Search
![search-example](https://github.com/devcblack/snx/assets/94953425/7ba0afa4-f7f2-4725-88be-4cba5914524b)
### Component-View
![component-view-example](https://github.com/devcblack/snx/assets/94953425/59b3e9a3-0325-4bca-9c2d-09de0214e6c3)
### Asset-Download
![asset-download-example](https://github.com/devcblack/snx/assets/94953425/b4afd2bf-77e3-45bc-bc0a-76f120b2e54c)


# Installation
### Clone Repository
```
git clone git@github.com:devcblack/snx.git
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
