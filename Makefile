.DEFAULT_GOAL = check

# Run all checks required to validate the codebase before merging.
.PHONY: check
check: test lint

# Run all tests in the project.
.PHONY: test
test:
	@go test ./... -race -count=1

# Lint the codebase.
.PHONY: lint
lint:
	@golangci-lint run
