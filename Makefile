.PHONY: ui run

binary_output = bin/gin-react-starter

all: ui
	go build -o $(binary_output) main.go

ui:
	cd ui && npm install && npm run build

run: all
	./$(binary_output)

