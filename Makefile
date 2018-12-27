.DEFAULT_GOAL := test

export GO111MODULE=on

clean:
	@echo "Clean ./bin"
	rm -rf bin pkg *.out

build: clean
	@echo "Build..."
	go build -o bin/terraform-provider-keycloak -tags netgo

install: build
	cp bin/* ~/.terraform.d/plugins/

test:
	go test ./... -v
vet: 
	go vet ./...

stest: vet
	go test ./... -short
