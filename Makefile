SERVER_NAME := cpctl
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOENVSET := $(GOCMD) env -w
BASEPATH := $(shell pwd)
BUILDDIR := $(BASEPATH)/dist

GOPROXY := http://10.8.22.212:8081/repository/go-group/,direct
GIT_NAME := gitlab.hycyg.com
GIT_IP := 10.8.22.102

# 编译参数
LDFLAGS := -w -s

# 交叉编译设置（编译Linux平台的静态二进制）
GOARCH := amd64
GOOS := linux
TARGET := $(BUILDDIR)/$(SERVER_NAME)

build_server_linux:
	@echo "开始构建$(SERVER_NAME)..."
	$(GOENVSET) GOSUMDB=off
	$(GOENVSET) GOPROXY=$(GOPROXY)
	$(GOENVSET) GOPRIVATE=$(GIT_NAME)
	$(GOENVSET) GOINSECURE=$(GIT_NAME)
	$(GOENVSET) GONOSUMDB=$(GIT_NAME)
	git config --global url."git@$(GIT_NAME):".insteadOf "http://$(GIT_NAME)"
	$(GOCMD) mod tidy
	CGO_ENABLED=0 GOARCH=$(GOARCH) GOOS=$(GOOS) $(GOBUILD) -ldflags "$(LDFLAGS)" -o $(TARGET)
