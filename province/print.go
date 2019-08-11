package province

import "fmt"

func Print() {
	for _, club := range Provinces {
		fmt.Println(club)
	}
}