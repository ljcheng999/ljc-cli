export GO111MODULE=on
# update app name. this is the name of binary
APP=ljc-cli
APP_EXECUTABLE="./bin/$(APP)"
ALL_PACKAGES=$(shell go list ./... | grep -v /vendor)
SHELL := /bin/bash # Use bash syntax

# Optional if you need DB and migration commands
# DB_HOST=$(shell cat config/application.yml | grep -m 1 -i HOST | cut -d ":" -f2)
# DB_NAME=$(shell cat config/application.yml | grep -w -i NAME  | cut -d ":" -f2)
# DB_USER=$(shell cat config/application.yml | grep -i USERNAME | cut -d ":" -f2)

# Optional colors to beautify output
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

## Quality
check-quality: ## runs code quality checks
	make lint
	make fmt
	make vet

# Append || true below if blocking local developement
lint: ## go linting. Update and use specific lint tool and options
	golangci-lint run --enable-all

vet: ## go vet
	go vet ./...

fmt: ## runs go formatter
	go fmt ./...

tidy: ## runs tidy to fix go.mod dependencies
	go mod tidy

## Test
test: ## runs tests and create generates coverage report
	make tidy
	make vendor
	go test -v -timeout 10m ./... -coverprofile=coverage.out -json > report.json

coverage: ## displays test coverage report in html mode
	make test
	go tool cover -html=coverage.out

## Build
build: ## build the go application
	mkdir -p bin/
	go build -ldflags="-w -s" -o $(APP_EXECUTABLE)
	@echo "Build passed"

run: ## runs the go binary. use additional options if required.
	make build
	chmod +x $(APP_EXECUTABLE)
	$(APP_EXECUTABLE)

###############################################################################################################
runh: ## runs the go binary. use additional options if required.
	make build
	chmod +x $(APP_EXECUTABLE)
	$(APP_EXECUTABLE) -h

dh:
	make build
	chmod +x $(APP_EXECUTABLE)
	$(APP_EXECUTABLE) deploy -h

d:
	make build
	chmod +x $(APP_EXECUTABLE)
	$(APP_EXECUTABLE) deploy \
		--git gitlab \
		--cloud-provider aws \
		--region us-east-1 \
		--role-arn arn:aws:iam::022985595394:role/CAPARole-local \
		--kubeconfig /tmp/kubeconfig \
		--chart-registry-url https://dachichang.github.io/helm-simple-web \
		--chart-version 1.1.1 \
		--chart simple-web \
		-c capi-cm-poc \
		-r jc-release \
		-n default \
		--log-json

change:
	export CUR="github.com/ljcheng999/ljc-cli"
	export NEW="github.com/ljcheng999/ljc-cli"
	go mod edit -module github.com/ljcheng999/ljc-cli
	find . -type f -name '*.go' -exec perl -pi -e 's/${CUR}/${NEW}/g' {} \;

###############################################################################################################

clean: ## cleans binary and other generated files
	go clean
	rm -rf bin/
	rm -f coverage*.out

vendor: ## all packages required to support builds and tests in the /vendor directory
	go mod vendor


wire: ## for wiring dependencies (update if using some other DI tool)
	wire ./...

# [Optional] mock generation via go generate
# generate_mocks:
# 	go generate -x `go list ./... | grep - v wire`

# [Optional] Database commands
## Database
migrate: build
	${APP_EXECUTABLE} migrate --config=config/application.test.yml

rollback: build
	${APP_EXECUTABLE} migrate --config=config/application.test.yml

.PHONY: all test build vendor
## All
all: ## runs setup, quality checks and builds
	make check-quality
	make test
	make build

.PHONY: help
## Help
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
