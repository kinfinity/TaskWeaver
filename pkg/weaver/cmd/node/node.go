/*
We want to have everything involved with making node commands
- instantiate
*/
package node

import "github.com/spf13/cobra"

func NewNodeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "node",
		Short:   "Manage nodes in the network",
		Long:    "",
		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			o, err := flags.ToOptions(f, cmd, baseName, args)
			//validate options
			// run cmd
			o.Run()
		},
	}
	return cmd
}

func (o *NodeOptions) Run() error {
	return nil
}
