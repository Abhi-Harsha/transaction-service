
configuration = ./setup/config-local.yaml

build:
	go build -o out/server ./cmd/api/main.go

run:
	./out/server -configFile=${configuration}

buildAndRun: build run