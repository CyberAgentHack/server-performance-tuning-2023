default: generate

.PHONY: test
test:
	go test -race -cover -parallel 4 ./...

.PHONY: run-local
run-local:
	@ENV_ENVIRONMENT=local go run main.go

.PHONY: run-cloud9
run-cloud9:
	@ENV_ENVIRONMENT=cloud9 go run main.go

.PHONY: generate
generate:
	go generate ${CURDIR}/...

.PHONY: setup
setup:
	sh scripts/setup.sh

## CICD用
ZIP_FILE_NAME=wsperf-application.zip

.PHONY: zip
zip:
	zip -r $(ZIP_FILE_NAME) . -x .git/\*
