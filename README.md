# nagios-check-rest [![Build Status](https://travis-ci.org/ynishi/nagios-check-rest.svg?branch=master)](https://travis-ci.org/ynishi/nagios-check-rest)

* nagios plugin check rest api written in go.
* This plugin is to check service availability from outside.

## Current Status

* Development

## Install

```bash
$ go get "github.com/ynishi/nagios-check-rest"
$ make
$ ls check_rest

# To install into $GOPATH/bin, do below.
$ make install

# To install to nagios plugin path, do below.
$ mv check_rest /path/to/nagios/plugin
```