## help: Print this help message.
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'


## web/install: Install js and css dependencies for the web.
web/install:
	npm --prefix ./web install

## web/server: Start a development server for the web.
web/server:
	npm --prefix ./web run dev

## web/build: Build js and css files for the web.
web/build:
	npm --prefix ./web run build

## web/lint: Lint js and css files and format for the web.
web/lint:
	npm --prefix ./web run lint && npm --prefix ./web run format