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

	humanChoice, _ := prompt.Run()
	computerChoice := getComputerChoice()

	fmt.Println("Computer Choice", computerChoice)

	fmt.Println(game.ScoreGame(humanChoice, computerChoice), "wins!")
}

func getComputerChoice() string {
	options := make([]string, 0)
	options = append(options,
		"rock",
		"paper",
		"scissors")

	rand.Seed(time.Now().Unix())
	return options[rand.Intn(len(options))]
}
