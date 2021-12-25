package main

import (
	"fmt"

	"github.com/kris10ansn/go-tic-tac-toe/src/cli"
	"github.com/kris10ansn/go-tic-tac-toe/src/game"
	"github.com/kris10ansn/go-tic-tac-toe/src/game/frontend/commandline"
)

func main() {
	var (
		frontEnd game.FrontEnd
		mode     = cli.GetArgument(1, "cli")
	)

	if mode == "cli" {
		frontEnd = commandline.New()
	} else {
		panic(fmt.Sprintf("Unsupported front-end mode: \"%s\" (arg 1)", mode))
	}

	game.PlayGame(frontEnd)
}
