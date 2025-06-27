GIT_USER_ID ?= taufiktriantono
GIT_REPO_ID ?= platform-sdk
GENERATOR ?= openapi-generator
GENERATOR_CMD = $(GENERATOR) generate
INPUT_DIR ?= api-specs
OUTPUT_DIR ?= gen
GENERATOR_ARGS = --git-user-id=$(GIT_USER_ID) \
                 --git-repo-id=$(GIT_REPO_ID) \
                 --additional-properties=packageName=client,enumClassPrefix=true

.PHONY: bundle bundle-docs clean generate $(APIS)

gen-client:
	$(GENERATOR_CMD) \
		-i openapi.yaml \
		-g go \
		-o $(OUTPUT_DIR) \
		$(GENERATOR_ARGS)

gen-gin-server:
	$(GENERATOR_CMD) \
		-i openapi.yaml \
		-g go-gin-server \
		-o $(OUTPUT_DIR) \
		$(GENERATOR_ARGS)

open:
	open docs/index.html

lint:
	redocly lint

bundle:
	redocly bundle

clean:
	rm -rf $(OUTPUT_DIR)

