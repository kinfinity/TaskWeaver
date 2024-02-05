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

		},
	}
	return cmd
}
