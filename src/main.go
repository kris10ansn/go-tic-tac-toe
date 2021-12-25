package main

import (
	"fmt"
	"os"

	"github.com/kris10ansn/go-tic-tac-toe/src/game"
	"github.com/kris10ansn/go-tic-tac-toe/src/game/frontend/commandline"
)

func main() {
	var (
		frontEnd game.FrontEnd
		mode     = GetCommandLineArgument(1, "cli")
	)

	if mode == "cli" {
		frontEnd = commandline.New()
	} else {
		panic(fmt.Sprintf("Unsupported front-end mode: \"%s\" (arg 1)", mode))
	}

	game.PlayGame(frontEnd)
}

func GetCommandLineArgument(index int, defaultValue string) string {
	if len(os.Args) >= index {
		return os.Args[index]
	} else {
		return defaultValue
	}
}
