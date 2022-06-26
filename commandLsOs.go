package main

import (
	"fmt"
	client "github.com/Virtomize/uii-go-api"
	"os"
	"text/tabwriter"
)

type lsOsCommand struct {
}

func (l *lsOsCommand) Run(ctx *context) error {
	c := client.NewClient(ctx.token)
	osList, err := c.ReadOperationSystems()
	if err != nil {
		return err
	}
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 8, 8, 1, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t", "Displayname", "Distribution", "Version", "Architecture")
	for _, osEntry := range osList {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t", osEntry.DisplayName, osEntry.Distribution, osEntry.Version, osEntry.Architecture)
	}
	fmt.Fprintf(w, "\n")

	return nil
}
