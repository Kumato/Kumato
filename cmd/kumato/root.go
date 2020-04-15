package main

import (
	"github.com/kumato/kumato/cmd/kumato/controller"
	"github.com/kumato/kumato/cmd/kumato/node"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{}
	cmd.AddCommand(controller.Cmd)
	cmd.AddCommand(node.Cmd)
	cmd.Execute()
}
