package cmd

import (
	"financas-cli/actions"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Atualiza um registro",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		actions.Update()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
