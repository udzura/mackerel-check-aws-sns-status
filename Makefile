VERSION := $(shell go run mackerel-check-aws-sns-status.go --version | sed 's/version: //')
.PHONY: install test setup default package release clean

default: test
	go build .

test: setup
	go test ./...

setup:
	go get ./...

package: test clean
	@GOOS=linux GOARCH=amd64 go build . && mkdir -p pkg && zip pkg/mackerel-check-aws-sns-status.linux-amd64-v$(VERSION).zip mackerel-check-aws-sns-status

release: package
	git push origin master
	ghr $(VERSION) pkg
	git fetch origin --tags

clean:
	rm -f mackerel-check-aws-sns-status
	find pkg -name '*.zip' | xargs rm
