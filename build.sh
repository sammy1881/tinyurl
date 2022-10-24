GOOS=darwin GOARCH=arm64 go build -o Mac/tinyurl_arm64
GOOS=darwin GOARCH=amd64 go build -o Mac/tinyurl_amd64
GOOS=windows GOARCH=amd64 go build -o Windows/tinyurl.exe
GOOS=linux GOARCH=amd64 go build -o Linux/tinyurl
GOOS=linux GOARCH=arm64 go build -o Linux/tinyurl_arm64