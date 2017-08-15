package main

import (
	"log"
	"reflect"
	"testing"

	"net/http"

	"net/url"

	"strings"

	"github.com/olorin/nagiosplugin"
)

var checkResult *nagiosplugin.Check
var optsResult *Opts

func init() {
	checkResult = nagiosplugin.NewCheck()
	checkResult.AddResult(nagiosplugin.OK, "good")

	optsResult = DefaultOpts
	u, err := url.Parse("http://localhost/data")
	if err != nil {
		log.Fatal(err)
	}
	optsResult.Url = u
	optsResult.Method = "POST"
	optsResult.Body = strings.NewReader(`{"message":"OK?"}`)

	SetOpts(optsResult)
}

func TestOpts(t *testing.T) {

	u, err := url.Parse("http://localhost/data")
	if err != nil {
		log.Fatal(err)
	}

	opts := &Opts{
		Client: http.DefaultClient,
		Url:    u,
		Method: "POST",
		Body:   strings.NewReader(`{"message":"OK?"}`),
	}
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
