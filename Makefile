.PHONY: test race gernerate lint fmt 

test: ## run tests
	go run github.com/onsi/ginkgo/v2/ginkgo --coverprofile=coverage.txt --covermode=atomic -cover ./...

race: ## run tests with race detector
	go run github.com/onsi/ginkgo/v2/ginkgo --race ./...

gernerate: ## run go generate
	go generate ./..

lint: ## run golangcli-lint checks
	go run honnef.co/go/tools/cmd/staticcheck ./...
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout 5m

fmt: ## gofmt and goimports all go files
	go run mvdan.cc/gofumpt -l -w -extra .
	find . -name '*.go' -exec go run golang.org/x/tools/cmd/goimports -w {} +