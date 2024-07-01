package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "navidad",
		Usage: "Aplicación para saber todo sobre la navidad...",
		Commands: []*cli.Command{
			{
				Name:    "semanas-para-navidad",
				Aliases: []string{"s"},
				Usage:   "Te muestra cuantas semanas faltan para navidad",
				Action: func(cCtx *cli.Context) error {
					DrawInTerminalWeeks()
					return nil
				},
			},
			{
				Name:    "es-navidad",
				Aliases: []string{"e"},
				Usage:   "Te dice si es navidad",
				Action: func(cCtx *cli.Context) error {
					DrawInTerminalChristmas()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
