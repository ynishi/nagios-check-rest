package main

import (
	"github.com/olorin/nagiosplugin"
)

func Check() (chk *nagiosplugin.Check, err error){
	check := nagiosplugin.NewCheck()
	defer check.Finish()
	check.AddResult(nagiosplugin.OK, "good")
	return check, nil
}

func main() {
	Check()
}
