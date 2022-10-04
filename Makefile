.PHONY: update
update:
	@go mod tidy \
		&& go mod vendor

.PHONY: run
run: update
	@go run ./cmd/api
