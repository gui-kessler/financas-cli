package cmd

import (
	"financas-cli/actions"

	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Realiza consultas das contas",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		actions.Query()
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}
