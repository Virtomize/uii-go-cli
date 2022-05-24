package main

import (
	"github.com/alecthomas/kong"
)

type commandStructure struct {
	Token string       `help:"The token used for authentication with the cloud-api." env:"CLOUDAPITOKEN" required:""`
	Build buildCommand `cmd:"" help:"Build a self installing iso"`
	Ls    struct {
		Os      lsOsCommand       `cmd:"" help:"List available operation systems"`
		Package lsPackagesCommand `cmd:"" help:"List available packages for a given operation systems"`
	} `cmd:"" help:"List objects"`
}

var cli commandStructure

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&context{token: cli.Token})
	ctx.FatalIfErrorf(err)
}
