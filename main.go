package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jones2026/go-tdd/game"
	"github.com/manifoldco/promptui"
)

func main() {
	fmt.Println("Let's play rock, paper, scissors!")

	prompt := promptui.Prompt{
		Label: "Human Choice ",
		// Validate: game.ValidateCompetitor,
	}

	human, _ := prompt.Run()

	competitors := make([]string, 0)
	competitors = append(competitors,
		"rock",
		"paper",
		"scissors")

	rand.Seed(time.Now().Unix())
	computer := competitors[rand.Intn(len(competitors))]
	fmt.Println("Computer chose", computer)

	fmt.Println(game.ScoreGame(human, computer), "wins!")
}
