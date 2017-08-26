# nagios-check-rest [![Build Status](https://travis-ci.org/ynishi/nagios-check-rest.svg?branch=master)](https://travis-ci.org/ynishi/nagios-check-rest)

* nagios plugin check rest api written in go.
* This plugin is to check service availability from outside.

## Current Status

* version 1.0
* basic features implemented. 

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

## Nagios setting

* nrpe command
```
command[check_rest]=/usr/lib64/nagios/plugins/check_rest -url $ARG1$ -w $ARG2$ -c $ARG3$ -h $ARG4$4-d $ARG5$
```
* nagios server
```
check_command check_nrpe!check_rest!http://localhost 20 5 'Authorization: apikey xxx' '{"message":"OK?"}'
```
