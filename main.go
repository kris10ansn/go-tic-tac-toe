package main

import (
	"os"

	"github.com/kris10ansn/go-tic-tac-toe/internal/players/cli"
	"github.com/kris10ansn/go-tic-tac-toe/internal/players/random"
	"github.com/kris10ansn/go-tic-tac-toe/pkg/game"
)

func main() {
	mode := getCommandLineArgument(1, "cli")

	if mode == "cli" {
		m := &cli.CommandlineManager{}

		game := game.CreateGame(
			m.CreateCommandlinePlayer(),
			m.CreateCommandlinePlayer(),
		)

		game.Play()
	}

	if mode == "cli-random" {
		cli := &cli.CommandlineManager{}

		game := game.CreateGame(
			cli.CreateCommandlinePlayer(),
			&random.RandomPlayer{},
		)

		game.Play()
	}
}

func getCommandLineArgument(index int, defaultValue string) string {
	if len(os.Args) > index {
		return os.Args[index]
	} else {
		return defaultValue
	}
}
