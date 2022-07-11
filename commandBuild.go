package main

import (
	client "github.com/Virtomize/uii-go-api"
)

type buildCommand struct {
	Output       string `help:"Output path to be written." required:"" arg:""`
	Distribution string `help:"The distribution of the iso. See ls os." required:"" arg:""`
	Version      string `help:"The version of the distribution." required:"" arg:""`
	Hostname     string `help:"The host name." required:"" arg:""`

	Password string   `help:"The root password." default:"virtomize"`
	Locale   string   `help:"The locale." default:"en_US"`
	Packages []string `help:"A list of packages to install."`
}

func (l *buildCommand) Run(ctx *context) error {
	c, err := client.NewClient(ctx.token)
	if err != nil {
		return err
	}

	err = c.Build(l.Output, client.BuildArgs{
		Distribution: l.Distribution,
		Version:      l.Version,
		Hostname:     l.Hostname,
		Networks: []client.NetworkArgs{
			{
				DHCP:       true,
				NoInternet: false,
			},
		},
	}, client.BuildOpts{
		Password: l.Password,
		Locale:   l.Locale,
		Packages: l.Packages,
	})
	return err
}
