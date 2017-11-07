# var
BIN = subscriber
CONFIG_FILE = config.json
DEPLOY_HOST = host
DEPLOY_BIN_DIRECTORY = /path/to/deploy
DEPLOY_CONFIG_DIRECTORY = /path/to/config

# target
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BIN) *.go

deploy-config:
	scp $(CONFIG_FILE).dist $(DEPLOY_HOST):$(DEPLOY_CONFIG_DIRECTORY)/$(CONFIG_FILE)

deploy-bin:
	scp $(BIN) $(DEPLOY_HOST):$(DEPLOY_BIN_DIRECTORY)/$(BIN)

deploy:
	$(MAKE) -s build-linux
	$(MAKE) -s deploy-bin
	$(MAKE) -s clean

run:
	go run *.go -config $(CONFIG_FILE)

clean:
	rm $(BIN)

code-quality:
	@echo "== GOLINT =="
	@find . -type d | xargs -L 1 golint
	@echo "== GO VET =="
	@find . -name "*.go" -exec go vet {} \;
