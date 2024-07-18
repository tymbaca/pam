package main

import "github.com/urfave/cli/v2"

type Generator interface {
	Generate(*cli.Context) error
}
