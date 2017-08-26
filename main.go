package main

import (
	"net/http"
	"time"

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

	time_start := time.Now()
	resp, err := options.Client.Do(options.Request)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var responseStatus *ResponseStatus
	switch {
	case time.Since(time_start) >= options.Warning:
		responseStatus = &ResponseStatus{
			Status:  nagiosplugin.WARNING,
			Message: "Warning: Response time overed",
		}
	case time.Since(time_start) >= options.Critical:
		responseStatus = &ResponseStatus{
			Status:  nagiosplugin.CRITICAL,
			Message: "Critical: Response time overed",
		}
	}
	if responseStatus != nil {
		check.AddResult(responseStatus.Status, responseStatus.Message)
		return check, nil
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	responseStatus, err = options.CheckFunc(buf)
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
	Warning   time.Duration
	Critical  time.Duration
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
		headerStr   string
		urlStr      string
		methodStr   string
		dataStr     string
		warningInt  int
		criticalInt int
	)

	f := flag.CommandLine
	f.StringVar(&urlStr, "url", "http://localhost", "url")
	f.StringVar(&urlStr, "u", "http://localhost", "short name of url")
	f.StringVar(&methodStr, "method", "GET", "method")
	f.StringVar(&methodStr, "m", "GET", "short name of method")
	f.IntVar(&warningInt, "warning", 5, "warning response time")
	f.IntVar(&warningInt, "w", 5, "short name of warning")
	f.IntVar(&criticalInt, "critical", 20, "critical response time")
	f.IntVar(&criticalInt, "c", 20, "short name of critical")
	f.StringVar(&headerStr, "header", "", "request header")
	f.StringVar(&headerStr, "h", "", "short name of header")
	f.StringVar(&dataStr, "data", "", "request body")
	f.StringVar(&dataStr, "d", "", "short name of data")

	f.Parse(args[1:])
	for _, arg := range f.Args() {
		if arg == "help" {
			f.Usage()
			return nil
		}
	}
	opts = DefaultOpts
	opts.Client.Timeout = time.Duration(warningInt+1) * time.Second
	opts.Warning = time.Duration(warningInt) * time.Second
	opts.Critical = time.Duration(criticalInt) * time.Second
	opts.Request, _ = http.NewRequest(methodStr, urlStr, strings.NewReader(dataStr))
	hs := strings.Split(headerStr, ":")
	if len(hs) > 1 {
		opts.Request.Header.Set(strings.Trim(hs[0], " "), strings.Trim(hs[1], " "))
	}
	return opts
}

func main() {
	opts := Parse(os.Args)
	if opts == nil {
		os.Exit(1)
	}
	SetOpts(opts)
	Check()
}
