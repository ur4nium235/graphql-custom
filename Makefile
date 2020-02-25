PROJECT_NAME := "project-base-go"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

# push default
default:
	@echo "==> Commit and push default <=="
	@ ./scripts/push_default.sh

test: ## Run unittests
	@go test -short ${PKG_LIST}

download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing tools from tools.go
#	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %
