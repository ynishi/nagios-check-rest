package main

import (
	"net/http"
	"net/url"

	"io"

	"github.com/olorin/nagiosplugin"
)

func Check() (chk *nagiosplugin.Check, err error) {
	check := nagiosplugin.NewCheck()
	defer check.Finish()

	check.AddResult(nagiosplugin.CRITICAL, "critical")
	return check, nil
}

type Opts struct {
	Client *http.Client
	Url    *url.URL
	Method string
	Body   io.Reader
}

var DefaultOpts = &Opts{
	Client: http.DefaultClient,
}

var options *Opts

func SetOpts(opts *Opts) {
	options = opts
}
func main() {
	Check()
}
