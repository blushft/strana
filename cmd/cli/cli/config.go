package cli

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

func configCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Print Strana Configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := configure()
			if err != nil {
				return err
			}

			spew.Dump(conf)

			return nil
		},
	}
}
