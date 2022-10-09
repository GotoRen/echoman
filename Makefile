DOCKER=docker
COMPOSE=docker compose
UP=$(COMPOSE) -f compose.yaml up -d
EXEC=$(COMPOSE) exec
LOGS=$(COMPOSE) logs -f



# docker-compose
#===============================================================
all: up

build: ## docker build
	$(BUILD)

up: ## docker up
	$(UP)

logs: $(LOGS) ## docker logs 

.PHONY: exec/server
exec/server: ## exec server container
	$(EXEC) server bash

.PHONY: exec/client
exec/client: ## exec client container
	$(EXEC) client bash

.PHONY: logs/server
logs/server: ## logs server container
	$(LOGS) server

.PHONY: logs/client
logs/client: ## logs client container
	$(LOGS) client

stop: ## docker stop
	$(COMPOSE) stop

down: ## docker down
	$(COMPOSE) down

down/all: ## delete images, network, containers
	$(DOCKER) system prune --all

down/vol: ## delete volumes
	$(DOCKER) volume prune



# Makefile config
#===============================================================
help: ## display this help screen
	@grep -E '^[a-zA-Z/_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
