package main

import "os"

func Example() {
	opts := Parse(os.Args)
	SetOpts(opts)
	Check()
	// Output:
	// OK: 10
}
