.PHONY: run docs

run: docs
	go run main.go

docs:
	swag init
