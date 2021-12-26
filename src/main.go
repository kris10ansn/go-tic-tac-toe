package main

import (
	"fmt"

	"github.com/kris10ansn/go-tic-tac-toe/src/cli"
	"github.com/kris10ansn/go-tic-tac-toe/src/game"
	"github.com/kris10ansn/go-tic-tac-toe/src/game/frontend/commandline"
	"github.com/kris10ansn/go-tic-tac-toe/src/game/frontend/web"
)

func main() {
	switch mode := cli.GetArgument(1, "cli"); mode {
	case "cli":
		game.PlayGame(commandline.New())
	case "web":
		gameServer := web.CreateServer()
		gameServer.Run()
	default:
		panic(fmt.Sprintf("Unsupported mode: \"%s\" (arg 1)", mode))
	}
}
