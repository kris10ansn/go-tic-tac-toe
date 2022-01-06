package main

import (
	"fmt"

	"github.com/kris10ansn/go-tic-tac-toe/internal/frontends/commandline"
	"github.com/kris10ansn/go-tic-tac-toe/internal/frontends/web"
	"github.com/kris10ansn/go-tic-tac-toe/pkg/cli"
	"github.com/kris10ansn/go-tic-tac-toe/pkg/game"
)

func main() {
	switch mode := cli.GetArgument(1, "cli"); mode {
	case "cli":
		game.PlayGame(commandline.New())
	case "web":
		server := web.CreateServer()
		server.Run()
	default:
		fmt.Printf("Unsupported mode: \"%s\" (arg 1), exiting...\n", mode)
	}
}
