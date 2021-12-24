package main

import (
	"fmt"
	"os"

	"github.com/kris10ansn/go-tic-tac-toe/src/cli"
	"github.com/kris10ansn/go-tic-tac-toe/src/game"
)

func main() {
	var noArgs bool = len(os.Args) <= 1
	var frontEnd game.FrontEnd

	if noArgs {
		fmt.Println("No command line arguments passed, running as command line interface")
	}

	if noArgs || os.Args[1] == "cli" {
		frontEnd = cli.New()
	} else {
		panic(fmt.Sprintf("Unsupported front-end mode: \"%s\" (arg 1)", os.Args[1]))
	}

	game.PlayGame(frontEnd)
}
