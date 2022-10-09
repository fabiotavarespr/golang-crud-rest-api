PWD = $(shell pwd -L)
GOCMD=go
GOTEST=$(GOCMD) test
DOCKER_CMD=docker
DOCKER_COMPOSE_CMD=docker-compose
PATH_DOCKER_COMPOSE_FILE=resources/scripts/docker-compose/docker-compose.yaml

ifeq ($(OS), Windows_NT)
	SHELL := powershell.exe
	.SHELLFLAGS := -NoProfile -Command
	SHELL_VERSION = $(shell (Get-Host | Select-Object Version | Format-Table -HideTableHeaders | Out-String).Trim())
	OS = $(shell "{0} {1}" -f "windows", (Get-ComputerInfo -Property OsVersion, OsArchitecture | Format-Table -HideTableHeaders | Out-String).Trim())
	PACKAGE = $(shell (Get-Content go.mod -head 1).Split(" ")[1])
	HELP_CMD = Select-String "^[a-zA-Z_-]+:.*?\#\# .*$$" "./Makefile" | Foreach-Object { $$_data = $$_.matches -split ":.*?\#\# "; $$obj = New-Object PSCustomObject; Add-Member -InputObject $$obj -NotePropertyName ('Command') -NotePropertyValue $$_data[0]; Add-Member -InputObject $$obj -NotePropertyName ('Description') -NotePropertyValue $$_data[1]; $$obj } | Format-Table -HideTableHeaders @{Expression={ $$e = [char]27; "$$e[36m$$($$_.Command)$${e}[0m" }}, Description
else
	SHELL := bash
	SHELL_VERSION = $(shell echo $$BASH_VERSION)
	UNAME := $(shell uname -s)
	VERSION_AND_ARCH = $(shell uname -rm)
	ifeq ($(UNAME),Darwin)
		OS = macos ${VERSION_AND_ARCH}
	else ifeq ($(UNAME),Linux)
		OS = linux ${VERSION_AND_ARCH}
	else
    $(error OS not supported by this Makefile)
	endif
	PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
	HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
endif

.DEFAULT_GOAL := help
.PHONY: help
help: ## Show this help
	@${HELP_CMD}

about: ## Display info related to the build
	@echo "OS: ${OS}"
	@echo "Shell: ${SHELL} ${SHELL_VERSION}"
	@echo "Go version: $(shell go version)"
	@echo "Go package: ${PACKAGE}"

tidy: ## Downloads go dependencies
	@$(GOCMD) mod tidy

fmt: tidy ## Run go fmt
	@$(GOCMD) fmt ./...

run-api: fmt ## Run api
	@$(GOCMD) run ./cmd/server/main.go

.PHONY: test
test: fmt clean-cache ## Launch tests
	@$(GOTEST) -race ./...

clean-cache: ## Clean cache
	@$(GOCMD) clean -testcache

docker-compose-up: ## Run docker-compose services of project
	$(DOCKER_COMPOSE_CMD) -f $(PATH_DOCKER_COMPOSE_FILE) up -d

docker-compose-down: ## Stop docker-compose services of project
	$(DOCKER_COMPOSE_CMD) -f $(PATH_DOCKER_COMPOSE_FILE) down --remove-orphans

docker-compose-restart: docker-compose-down docker-compose-up ## Restart docker-compose services of project

docker-compose-logs: ## Logs docker-compose containers of project
	$(DOCKER_COMPOSE_CMD) -f $(PATH_DOCKER_COMPOSE_FILE) logs -f

docker-compose-ps: ## List docker-compose containers of project
	$(DOCKER_COMPOSE_CMD) -f $(PATH_DOCKER_COMPOSE_FILE) ps