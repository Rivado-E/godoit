package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type Task struct {
	title  string
	status bool
	desc   string
	id     int
}

func main() {
	var tasks []Task
	var language string
	newTask := false

	app := &cli.App{
		Name:  "goDoIt",
		Usage: "get it done!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				Value:       "english",
				Usage:       "language for the greeting",
				Destination: &language,
			},
			&cli.BoolFlag{
				Name:    "view",
				Aliases: []string{"v"},
			},
			&cli.BoolFlag{
				Name:        "new",
				Aliases:     []string{"n"},
				Usage:       "Create a new task",
				Destination: &newTask,
			},
		},
		Action: func(cCtx *cli.Context) error {
			name := "someone"
			if cCtx.NArg() > 0 {
				name = cCtx.Args().Get(0)
			}
			if language == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
		// Action: func(cCtx *cli.Context) error {
		// 	var name string
		// 	var desc string
		// 	// var status bool
		// 	if newTask {
		// 		fmt.Scanf("Name: ", &name)
		// 		fmt.Scanf("Description: ", &desc)
		// 	}
		// 	return nil
		//
		// },
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
