build_server_linux:
	@echo "开始构建$(SERVER_NAME)..."
	@echo "修改hosts配置，加入$(GIT_NAME)"
	echo "10.8.22.102 gitlab.hycyg.com" >> /etc/hosts
	go env -w GOSUMDB=off
	go env -w GOPROXY=$(GOPROXY)
	go env -w GOPRIVATE=$(GIT_NAME)
	go env -w GOINSECURE=$(GIT_NAME)
	go env -w GONOSUMDB=$(GIT_NAME)
	git config --global url."git@$(GIT_NAME):".insteadOf "http://$(GIT_NAME)"
	go mod tidy
	CGO_ENABLED=0 GOARCH=$(GOARCH) GOOS=$(GOOS) go build -ldflags "-X main.version=$(VERSION)"  -o $(SERVER_NAME)


kubernetes_deploy:
	@echo "正在使用 helm 部署 $(SERVER_NAME) $(VERSION) ..."
	@echo "环境变量 SERVICE_PORT: $(SERVICE_PORT) NACOS_URL: $(NACOS_URL)"
	@echo "CHART_FILE: $(CHART_FILE)"
	@echo "NAMESPACE: $(NAMESPACE)"
	@echo "KUBECONFIG_PATH: $(KUBECONFIG_PATH)"

	@echo "Checking if Helm repo exists..."
	if ! helm repo list | grep -q "paas"; then \
            echo "Adding Helm repo..."; \
            helm repo add paas $(CHART_REPO); \
        else \
            echo "Helm repo already exists, skipping repo add."; \
        fi
	helm repo update
	helm upgrade --install $(SERVER_NAME) $(CHART_FILE) -n $(NAMESPACE) \
	--set paas.name=$(SERVER_NAME),\
	paas.tag=$(VERSION),\
	paas.service.port=$(SERVICE_PORT),\
	paas.env[0].name=NACOS_URL,paas.env[0].value=$(NACOS_URL),paas.env[1].name=NACOS_GROUP,paas.env[1].value=paas,\
	fullnameOverride=$(SERVER_NAME) \
	--kube-token $(KUBE_TOKEN) --kube-apiserver $(KUBE_APISERVER) --kube-insecure-skip-tls-verify