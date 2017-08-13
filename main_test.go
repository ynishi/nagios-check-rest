package main

import (
	"reflect"
	"testing"

	"github.com/olorin/nagiosplugin"
)

var checkResult *nagiosplugin.Check

func init() {
	checkResult = nagiosplugin.NewCheck()
	checkResult.AddResult(nagiosplugin.OK, "good")
}

func TestCheck(t *testing.T) {

	check, err := Check()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(checkResult, check) {
		t.Fatalf("Didn't match check Result.\n want: %q,\n have: %q", checkResult, check)
	}
}
