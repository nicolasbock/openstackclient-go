openstack-go: *.go go.mod cmd/*.go lib/*.go
	go build -v -o $@ main.go

.PHONY: test
test:
	go test -v ./...
