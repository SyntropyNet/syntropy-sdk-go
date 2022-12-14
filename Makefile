DOCKERCMD=$(shell which docker)
OPENAPI_GENERATOR_VERSION=latest
OPENAPI := $(DOCKERCMD) run --rm -it -v "${PWD}/api:/input" -v "${PWD}/syntropy:/output" openapitools/openapi-generator-cli:$(OPENAPI_GENERATOR_VERSION)

SYNTROPY_OPENAPI_SPEC=api/openapi.yaml
SYNTROPY_OPENAPI_URL=https://api.syntropystack.com/docs/openapi.yaml

download:
	wget ${SYNTROPY_OPENAPI_URL} -O ${SYNTROPY_OPENAPI_SPEC}

validate:
	@$(OPENAPI) validate -i /input/openapi.yaml

generate:
	rm -rf syntropy/*
	@$(OPENAPI) generate -i /input/openapi.yaml -g go -o /output -p enumClassPrefix=true --package-name syntropy
	gofmt -w syntropy

update: download generate