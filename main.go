package main

import (
	"os"

	"github.com/tymbaca/gorange/cli/generators/gogen"
	"github.com/urfave/cli/v2"
	// "github.com/urfave/cli/v2"
)

func main() {
	goGen := gogen.New()

	app := &cli.App{
		Name:  "pam",
		Usage: "generate project layout",
		// Action: func(ctx *cli.Context) error {
		// 	fmt.Println(ctx.NArg())
		// 	fmt.Println(ctx.Args().Get(0))

		// 	return nil
		// },
		Commands: cli.Commands{
			{
				Name:    "golang",
				Aliases: []string{"go"},
				Usage:   "generate go project template",
				Action: func(ctx *cli.Context) error {
					return goGen.Generate(ctx)
				},
			},
		},
	}

	fatal(app.Run(os.Args))
}

func fatal(err error) {
	if err != nil {
		panic(err)
	}
}
