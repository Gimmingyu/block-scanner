GOCMD=go

CMD = cmd
INIT = init
ENTRY = main.go

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

ETHEREUM = ethereum

all: help

## Version:
version: ## Show version.
	@$(GOCMD) version

## Ethereum:
ethereum: ## Build ethereum.
	@echo "${GREEN}Building ${CYAN}$(ETHEREUM)${GREEN}...${RESET}"
	@$(GOCMD) build -o $(INIT)/$(ETHEREUM) $(CMD)/$(ETHEREUM)/$(ENTRY)

## Run ethereum:
run-ethereum: ethereum ## Run ethereum.
	@echo "${GREEN}Running ${CYAN}$(ETHEREUM)${GREEN}...${RESET}"
	@$(GOCMD) run $(CMD)/$(ETHEREUM)/$(ENTRY)

## Help:
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

.PHONY: version help