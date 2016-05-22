precommit: generate fmt license test

test:
	@go test -race $(shell go list ./... | grep -v vendor)
fmt:
	@go fmt $(shell go list ./... | grep -v vendor | grep -v testdata)
license:
	@go run $(GOPATH)/src/github.com/limetext/tasks/gen_license.go
generate:
	@go generate $(shell go list ./... | grep -v /vendor/)
fast_test:
	@go test $(shell go list ./... | grep -v vendor)

tasks:
	go get -d -u github.com/limetext/tasks
