package main

import (
	"fmt"
	client "github.com/Virtomize/uii-go-api"
)

type lsPackagesCommand struct {
	Distribution string `help:"The name of the distribution for which the packages should be listed." required:"" arg:""`
	Version      string `help:"The version of the distribution for which the packages should be listed." required:"" arg:""`
	Architecture string `help:"The architecture (32/64) of the distribution for which the packages should be listed." optional:"" default:"" arg:""`
}

func (l *lsPackagesCommand) Run(ctx *context) error {
	c := client.NewClient(ctx.token)
	packageList, err := c.ReadPackages(client.PackageArgs{
		Distribution: l.Distribution,
		Version:      l.Version,
	}, client.PackageOpts{Architecture: l.Architecture})
	if err != nil {
		return err
	}

	for _, p := range packageList {
		fmt.Printf("%s\n", p)
	}
	return nil
}
