package main

import (
	"fmt"
	"os"

	"github.com/kris10ansn/go-tic-tac-toe/internal/frontends/cli"
	"github.com/kris10ansn/go-tic-tac-toe/internal/frontends/web"
	"github.com/kris10ansn/go-tic-tac-toe/pkg/game"
)

func main() {
	switch mode := getCommandLineArgument(1, "cli"); mode {
	case "cli":
		game.PlayGame(cli.New())
	case "web":
		server := web.CreateServer()
		server.Run()
	default:
		fmt.Printf("Unsupported mode: \"%s\" (arg 1), exiting...\n", mode)
	}
}

func getCommandLineArgument(index int, defaultValue string) string {
	if len(os.Args) > index {
		return os.Args[index]
	} else {
		return defaultValue
	}
}
