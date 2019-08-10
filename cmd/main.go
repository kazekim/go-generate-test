package main

import "github.com/kazekim/go-generate-test/premierleague"

var VersionString = "unset"

//go:generate go run ../gen.go

func main() {
	premierleague.PrintFootballClubs()
}
