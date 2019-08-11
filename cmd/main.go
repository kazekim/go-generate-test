package main

import "github.com/kazekim/go-generate-test/province"

var VersionString = "unset"

//go:generate go run gen.go

func main() {
	province.Print()
}
