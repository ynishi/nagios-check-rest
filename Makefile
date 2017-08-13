#
# Makefile to build and install check_rest
# (to change binary name)
#

.PHONY: build test clean install

build: clean test
	go build -i -o check_rest

clean:
	rm -f check_rest

test:
	go test

install:
	mv check_rest ${GOPATH}/bin/.
