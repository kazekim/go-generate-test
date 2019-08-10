package premierleague

import "fmt"

//go:generate go run gen.go

func PrintFootballClubs() {
	for _, club := range FootballClubs {
		fmt.Println(club)
	}
}