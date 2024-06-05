include .envrc

.PHONY: run
run:
	@go run ./cmd -username=${USERNAME} -oauth=${OAUTH}