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
	echo "test"
