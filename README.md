# microsoft-wallpapers

## Preparation
Prepare an input file that contains a list of links from the Microsoft website that lead to wallpapers. This list is provided in `wallpapers.txt`, but may be updated via
```
lynx -dump -listonly https://support.microsoft.com/en-us/windows/wallpapers-5cfa0cc7-b75a-165a-467b-c95abaf5dc2a\#WindowsVersion\=Windows_10 | grep kbdevstorage > wallpapers.txt
```

## Running
To download wallpapers, run `main.go`, e.g.,
```sh
go run main.go
```
This may take a while. test
