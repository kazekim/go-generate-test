package province

import "fmt"

//go:generate go run gen.go

func Print() {
	for _, club := range Provinces {
		fmt.Println(club)
	}
}