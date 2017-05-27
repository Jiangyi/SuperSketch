.PHONY: build run view view-dev deps

BINARY=supersketch

build:
	go build -race -o ${BINARY}

run: build
	@echo "Running server..."
	./${BINARY}

clean:
	@echo "Removing compiled go binary"
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

get-glide:
	sudo add-apt-repository ppa:masterminds/glide && sudo apt-get update
	sudo apt-get install glide

deps:
	glide install
	cd ./view/; npm install

view:
	cd ./view/;	npm run build

view-dev:
	cd ./view/; npm run dev