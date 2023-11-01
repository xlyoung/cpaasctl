

build_server_linux:
	@echo "开始构建$(SERVER_NAME)..."
	go env -w GOSUMDB=off
	go env -w GOPROXY=$(GOPROXY)
	go env -w GOPRIVATE=$(GIT_NAME)
	go env -w GOINSECURE=$(GIT_NAME)
	go env -w GONOSUMDB=$(GIT_NAME)
	git config --global url."git@$(GIT_NAME):".insteadOf "http://$(GIT_NAME)"
	go mod tidy
	CGO_ENABLED=0 GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o $(TARGET)
