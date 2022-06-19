package main

import (
	"github.com/alecthomas/kong"
)

type commandStructure struct {
	Token string       `help:"The token used for authentication" env:"UIIAPITOKEN" required:""`
	Build buildCommand `cmd:"" help:"Builds a unattended installing iso"`
	Ls    struct {
		Os      lsOsCommand       `cmd:"" help:"Lists the available operation systems"`
		Package lsPackagesCommand `cmd:"" help:"Lists the available packages for a given operation systems"`
	} `cmd:"" help:"List objects"`
}

var cli commandStructure

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&context{token: cli.Token})
	ctx.FatalIfErrorf(err)
}
