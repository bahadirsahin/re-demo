# Makefile for 're-demo'
.PHONY: all

all: help

## Help
help: ## 
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    %-20s%s\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  \n%s\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)

## Test
test: ## 
	make -i clean
	make -i build
	make -i run
	make -i logs

## Build
build: ## 
	cd re-demo-api; docker build -f Dockerfile -t re-demo-api-img:latest .

## Clean
clean: ## 
	docker compose -p re rm -fs re-demo-api-svc

## Deploy
deploy: ## 
	make -i build
	make -i run
	
## Logs
logs: ## 
	docker compose -p re logs -f re-demo-api-svc

## Run
run: ## 
	docker compose -p re up -d re-demo-api-svc
