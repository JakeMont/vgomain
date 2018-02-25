package main

import (
	"github.com/jakemont/vgotest"
	"github.com/jakemont/vgotest/foo"
)

func main() {
	println("vtest", vgotest.Version, "foo", foo.Version)
}
