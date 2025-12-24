default: generate

.PHONY: test
test:
	go test -race -cover -parallel 4 ./...

.PHONY: run
run:
	@ENV_ENVIRONMENT=local go run main.go

.PHONY: generate
generate:
	go generate ${CURDIR}/...

.PHONY: setup
setup:
	sh scripts/setup.sh
