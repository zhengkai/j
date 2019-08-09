SHELL:=/bin/bash

coverage = coverage.out

default:
	@rm -rf log-dir
	@rm -f log-file
	@rm -f "$(coverage)"
	@go test -v -test.bench=".*" -test.benchmem -coverpkg github.com/zhengkai/j -covermode=count "-coverprofile=$(coverage)"
	@if [ -n "$$TRAVIS_JOB_ID" ]; then \
		"$$HOME/gopath/bin/goveralls" "-coverprofile=$(coverage)" -service=travis-ci -repotoken "$$COVERALLS_TOKEN"; \
	fi
