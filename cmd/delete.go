package cmd

import (
	"financas-cli/actions"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove um registro",
	Long:  `Remove um registro de despesa ou receita.`,

	Run: func(cmd *cobra.Command, args []string) {
		actions.Delete()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
