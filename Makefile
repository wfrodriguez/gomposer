TAGLINE := "Simple aplicación CLI para la gestión básica de bases de datos"

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RED    := $(shell tput -Txterm setaf 1)
RESET  := $(shell tput -Txterm sgr0)

# Nombre del binario
BINNAME = gql
# Ubicación del archivo main
MAIN = cmd/gql/main.go
# Argumentos de la línea de comandos
ARGS = -k select -l

VERSION := 0.1.0
COMMIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_TIMESTAMP := $(shell date '+%Y-%m-%dT%H:%M:%S')

TARGET_MAX_CHAR_NUM := 20

.DEFAULT_GOAL := help
.PHONY: help build validate run-go run-bin



## Valida si los programas obligatorios se encuentran instalados en el sistema
validate:
	@command -v go >/dev/null 2>&1 || { echo "${RED}Requiero el programa 'go' pero no está instalado. Abortando.${RESET}" >&2; exit 1; }

## Muestra este mensaje de ayuda
help:
	@echo ${TAGLINE}
	@echo ''
	@echo 'Modo de uso:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

## Compila la aplicación y genera un binario en ./dist
build: validate
	@mkdir -p ./dist
	go build -ldflags="-X main.Version='$(VERSION)' -X main.CommitHash='$(COMMIT_HASH)' -X main.BuildTimestamp='$(BUILD_TIMESTAMP)'" -o ./dist/$(BINNAME) $(MAIN)

## Inicia la aplicación desde el código fuente
run-go: validate
	@go run $(MAIN) $(ARGS)

## Ejecuta el binario generado
run-bin: build
	@./dist/$(BINNAME) $(ARGS)
