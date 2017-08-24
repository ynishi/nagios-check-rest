package main

import (
	"net/http"

	"encoding/json"

	"flag"
	"os"
	"strings"

	"io/ioutil"

	"github.com/olorin/nagiosplugin"
)

func Check() (chk *nagiosplugin.Check, err error) {
	check := nagiosplugin.NewCheck()
	defer check.Finish()

	resp, err := options.Client.Do(options.Request)
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	responseStatus, err := options.CheckFunc(buf)
	if err != nil {
		return nil, err
	}
	check.AddResult(responseStatus.Status, responseStatus.Message)
	return check, nil
}

type Opts struct {
	Client    *http.Client
	Request   *http.Request
	CheckFunc func([]byte) (responseStatus *ResponseStatus, err error)
}

type ResponseStatus struct {
	Status  nagiosplugin.Status `json:status`
	Message string              `json:message`
}

func DefalutCheckFunc(buf []byte) (resposneStatus *ResponseStatus, err error) {
	resposneStatus = &ResponseStatus{}
	err = json.Unmarshal(buf, resposneStatus)
	if err != nil {
		return nil, err
	}
	return resposneStatus, nil
}

var DefaultOpts = &Opts{
	Client:    http.DefaultClient,
	CheckFunc: DefalutCheckFunc,
}

var options *Opts

func SetOpts(opts *Opts) {
	options = opts
}

func Parse(args []string) (opts *Opts) {
	var (
		headerStr string
		urlStr    string
		methodStr string
		dataStr   string
	)

	f := flag.CommandLine
	f.StringVar(&headerStr, "h", "", "header")
	f.StringVar(&urlStr, "u", "", "url")
	f.StringVar(&methodStr, "m", "", "method")
	f.StringVar(&dataStr, "d", "", "data")

	f.Parse(args[1:])
	opts = DefaultOpts
	opts.Request, _ = http.NewRequest(methodStr, urlStr, strings.NewReader(dataStr))
	hs := strings.Split(headerStr, ":")
	opts.Request.Header.Set(strings.Trim(hs[0], " "), strings.Trim(hs[1], " "))
	return opts
}

func main() {
	SetOpts(Parse(os.Args))
	Check()
}
