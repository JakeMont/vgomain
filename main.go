package main

import (
	"github.com/JakeMont/vgotest"
	"github.com/JakeMont/vgotest/foo"
)

func main() {
	println("vtest", vgotest.Version, "foo", foo.Version)
}
