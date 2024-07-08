build-win:
	@GOOS=windows GOARCH=amd64 go build -o bin/windows/yoo .

build-linux:
	@GOOS=linux GOARCH=amd64 go build -o bin/linux/yoo .

build-mac:
	@GOOS=darwin GOARCH=amd64 go build -o bin/macos/yoo .
