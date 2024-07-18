package gogen

import (
	"errors"

	"github.com/urfave/cli/v2"
)

type gen struct {
}

func New() *gen {
	return &gen{}
}

func (g *gen) Generate(*cli.Context) error {
	return errors.New("not implemented")
}
