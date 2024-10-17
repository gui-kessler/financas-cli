package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Realiza consultas das contas",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("query called")
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}
