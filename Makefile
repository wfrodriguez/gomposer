# Nombre del proyecto
NAME = gomposer
# Nombre del binario
BINNAME = gps
# We can use such syntax to get main.go and other root Go files.
GO_FILES = $(wildcard *.go)
# Working Directory
WD = $(pwd)
# Otros
MAIN = cmd/gomposer/main.go
BIN = bin
VERSION = 0.1.0
COMMIT_HASH = $(git rev-parse --short HEAD)
BUILD_TIMESTAMP = $(date '+%Y-%m-%dT%H:%M:%S')
ARGS = post

# Valida si ciertos programas se encuentran instalados en el sistema
validate:
	@command -v reflex >/dev/null 2>&1 || { echo "Se requiere el programa 'reflex' pero no está instalado. Abortando." >&2; exit 1; }

help: # \tMuestra la ayuda de las diferentes tareas
	@echo "Tareas del proyecto $(NAME):"
	@echo
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m  $$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2 -d'#')\n"; done

run: # \t\tLa tarea 'run' ejecutará la aplicación desde el archivo MAIN.
	@go run $(MAIN) $(ARGS)

build: generate # \tLa tarea 'build' Compila la aplicación y genera un binario en ./$(BIN)
	@mkdir -p ./$(BIN)
	@go build -ldflags="-X main.Version='$(VERSION) ($(COMMIT_HASH) $(BUILD_TIMESTAMP))'" -o $(BIN)/$(BINNAME) $(MAIN)

# .PHONY is used for reserving tasks words
.PHONY:

