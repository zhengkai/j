SHELL:=/bin/bash

default:
	go test -v -test.bench=".*" -test.benchmem -coverpkg github.com/zhengkai/j -covermode=count -coverprofile=coverage.out
