# justserve
A super simple tool to serve static files

### Build

- Download the repo. Then go the directory. 
- `go build -o serve main.go` or `go install`

### Usage

`./justserve path/to/the/directory/to/be/served` (the default port number `9102` will be used)

`./justserve -p 8888 path/to/the/directory/to/be/served` (using the provided port number)

`./justserve path/to/the/directory/to/be/served -p 8888` (using the provided port number)