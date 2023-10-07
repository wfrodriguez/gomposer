.DEFAULT_GOAL := help
.PHONY: help validate build serve clean copy mkdir install

Tagline := "Gomposer: Conjunto de herramientas para generar páginas web estáticas"
TargetMaxCharNum := 15

Green  := $(shell tput -Txterm setaf 2)
Yellow := $(shell tput -Txterm setaf 3)
White  := $(shell tput -Txterm setaf 7)
Red    := $(shell tput -Txterm setaf 1)
Reset  := $(shell tput -Txterm sgr0)

Name := Mi proyecto
DistDir := dist
BuildDir := build
ServerPort := 3000
HtmlFiles := $(patsubst post/%.md,$(DistDir)/post/%/index.html,$(wildcard post/*.md))
HtmlTags := $(patsubst $(BuildDir)/tag/%.md,$(DistDir)/tag/%/index.html,$(wildcard $(BuildDir)/tag/*.md))
Rm := @/bin/rm

Pandoc := @pandoc
PandocOptions := --read markdown --write html5 --standalone  --listings
PandocPost := --table-of-contents --template template/post.html
PandocTag := --template template/tag.html
PandocIndex := --template template/index.html

# https://github.com/m3ng9i/ran
Ran := @ran
RanOptions := -root $(DistDir) -port $(ServerPort) -no-cache

## Instala las dependencias necesarias del proyecto (solo root puede hacer esto y unicamente funciona en SO Unix)
install:
ifneq ($(shell id -u), 0)
	@echo "$(Red)Debes ser root para realizar esta acción.$(Reset)"
else
	@echo "TODO: The action when the user is root here..."
endif

$(DistDir)/post/%/index.html : post/%.md # genera las páginas HTML
	@mkdir -p $(@D)
	$(Pandoc) $(PandocOptions) $(PandocPost) --output $(@D)/index.html $<
	@printf "󰒔 %s -> %s\n" $< "$(@D)/$(@F)"

$(DistDir)/tag/%/index.html : $(BuildDir)/tag/%.md # genera las páginas HTML
	@mkdir -p $(@D)
	$(Pandoc) $(PandocOptions) $(PandocTag) --output $(@D)/index.html $<
	@printf "󰒔 %s -> %s\n" $< "$(@D)/$(@F)"

## Muestra este mensaje de ayuda
help:
	@echo ${Tagline}
	@echo ''
	@echo '${Yellow}Proyecto:${Reset} ${Name}'
	@echo ''
	@echo 'Modo de uso:'
	@echo '  ${Yellow}make${Reset} ${Green}<target>${Reset}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${Yellow}%-$(TargetMaxCharNum)s${Reset} ${Green}%s${Reset}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST) | sort

## Valida si los programas necesarios se encuentran instalados en el sistema
validate:
	@command -v pandoc >/dev/null 2>&1 || { echo "Se requiere el programa 'pandoc' pero no está instalado. Es necesario que se instale para funcionar" >&2; exit 2; }
	@command -v ran >/dev/null 2>&1 || { echo "Se requiere el programa 'ran' pero no está instalado. Es necesario que se instale para funcionar" >&2; exit 2; }

## Genera los directorios necesarios
mkdir:
	@mkdir -p $(DistDir)/{post,tag,build,static/{css,js,img,fonts,html}}

## Borra los archivos temporales y los directorios
clean:
	 @echo "Cleaning..."
	 $(Rm) -rf $(DistDir)

## Genera los archivos indices generados por gomposer
index: mkdir
	@echo "Generando los índices..."
	@gomposer index
	$(Pandoc) $(PandocOptions) $(PandocTag)  --output $(DistDir)/post/index.html build/posts.md
	$(Pandoc) $(PandocOptions) $(PandocTag) --output $(DistDir)/tag/index.html build/tags.md
	$(Pandoc) $(PandocOptions) $(PandocIndex) --metadata title="$(Name)" --output $(DistDir)/index.html index.md


## copia los archivos necesarios (static/..., archivos necesarios)
copy:
	@echo "Copiando..."
	@cp -r static $(DistDir)

## Ejecuta las páginas HTML en un servidor web temporal (requiere el programa 'ran')
serve:
	@echo "Iniciando el servidor en el puerto $(ServerPort)..."
	$(Ran) $(RanOptions)

## Genera todos los archivos html
build: clean validate index $(HtmlFiles) $(HtmlTags) copy

all : build serve
