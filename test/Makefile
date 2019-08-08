SHELL:=/bin/bash

default:
	go test -v -test.bench=".*" -test.benchmem -coverpkg github.com/zhengkai/j -covermode=count -coverprofile=coverage.out
	if [ -n "$$TRAVIS_JOB_ID" ]; then \
		"$$HOME/gopath/bin/goveralls" -coverprofile=coverage.out -service=travis-ci -repotoken "$$COVERALLS_TOKEN"; \
	fi
