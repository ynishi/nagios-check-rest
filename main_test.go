package main

import (
	"log"
	"reflect"
	"testing"

	"net/http"

	"strings"

	"encoding/json"

	"io/ioutil"

	"github.com/olorin/nagiosplugin"
)

var checkResult *nagiosplugin.Check
var optsResult *Opts

const responseJson = `{
"status":10000,
"message":"OK"
}`

var optsFunc = func(buf []byte) (responseStatus *ResponseStatus, err error) {
	responseStatusOrig := &ResponseStatus{}
	err = json.Unmarshal(buf, responseStatusOrig)
	if err != nil {
		return nil, err
	}
	responseStatus = &ResponseStatus{
		Message: responseStatusOrig.Message,
	}
	switch responseStatusOrig.Status {
	case 10000:
		responseStatus.Status = nagiosplugin.OK
	case 20000:
		responseStatus.Status = nagiosplugin.WARNING
	case 30000:
		responseStatus.Status = nagiosplugin.CRITICAL
	default:
		responseStatus.Status = nagiosplugin.UNKNOWN
	}
	return responseStatus, nil
}

var request *http.Request

func init() {
	checkResult = nagiosplugin.NewCheck()
	checkResult.AddResult(nagiosplugin.OK, "good")

	var err error
	request, err = http.NewRequest("POST", "http://localhost/data/check", strings.NewReader(`{"message":"OK?"}`))
	if err != nil {
		log.Fatal(err)
	}
	optsResult = DefaultOpts
	optsResult.Request = request
	optsResult.CheckFunc = optsFunc

	SetOpts(optsResult)
}

func TestArgs(t *testing.T) {

	args := strings.Split(`check_rest -u http://localhost/data/check -m POST -h "Authorization: apikey xxx" -d "{\"message\":\"OK?\"}"`, " ")
	opts := Parse(args[1:])

	url := "http://localhost/data/check"
	if opts.Request.URL.String() != url {
		t.Errorf("failed parse URL.\n want: %q,\n have: %q\n", url, opts.Request.URL)
	}
	auth := []string{"apikey xxx"}
	if !reflect.DeepEqual(opts.Request.Header["Authorization"], auth) {
		t.Errorf("failed parse apikey.\n want: %q,\n have: %q\n", auth, opts.Request.Header["Authorization"])
	}

	r := opts.Request.Body
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	data := []byte(`{"message":"OK?"}`)
	if !reflect.DeepEqual(data, buf) {
		t.Errorf("failed parse body.\n want: %q,\n have: %q\n", data, buf)
	}

}

func TestOpts(t *testing.T) {

	opts := DefaultOpts
	opts.Request = request
	opts.CheckFunc = optsFunc

	if !reflect.DeepEqual(optsResult, opts) {
		t.Fatalf("Didn't match Opts.\n want: %q,\n have: %q\n", optsResult, opts)
	}
}

func TestCheck(t *testing.T) {

	check, err := Check()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(checkResult, check) {
		t.Fatalf("Didn't match check Result.\n want: %q,\n have: %q\n", checkResult, check)
	}
}
