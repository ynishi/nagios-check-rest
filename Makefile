#
# Makefile to build and install check_rest
# (to change binary name)
#

build:
	go build -i -o check_rest
install:
	mv check_rest ${GOPATH}/bin/.
