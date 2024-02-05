package wcmd

import (
	"os"

	"github.com/spf13/cobra"
)

type WeaverOptions struct {
	Arguments   []string
	ConfigFlags WeaverFlags
}

type WeaverFlags struct {
	//
}

type CliCommand interface {
	Validate() error
}

func NewWeaverCommand() *cobra.Command {
	wo := WeaverOptions{
		Arguments: os.Args,
	}
	return NewWeaverCommandWithArgs(wo)

}

func NewWeaverCommandWithArgs(o WeaverOptions) *cobra.Command {
	cmd := &cobra.Command{}

	return cmd
}

func (wo *WeaverOptions) Validate() error {
	//
	return nil
}
