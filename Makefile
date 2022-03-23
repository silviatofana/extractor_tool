generated: build

test:
	@go test ./...

build: test
	export CGO_ENABLED=0 export GOOS=linux && go build -a -tags netgo -ldflags '-w -X main.version=v1.1.0' -o extractor_tool_linux
	export CGO_ENABLED=0 export GOOS=darwin && go build -a -tags netgo -ldflags '-w -X main.version=v1.1.0' -o extractor_tool_osx
	export CGO_ENABLED=0 export GOOS=windows && go build -a -tags netgo -ldflags '-w -X main.version=v1.1.0' -o extractor_tool_windows.exe
	export GOOS=$GOOS_OLD