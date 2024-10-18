package cmd

import (
	"financas-cli/actions"

	"github.com/spf13/cobra"
)

var isCredit, isDebit bool
var name string
var value float64

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Cria um novo registro",
	Long:  `Pode ser criado registro de despesa ou receita.`,

	Run: func(cmd *cobra.Command, args []string) {
		if isCredit && isDebit {
			panic("Não é possível criar um registro de despesa e receita ao mesmo tempo")
		}

		if !isCredit && !isDebit {
			panic("É necessário informar o tipo de registro a ser criado")
		}

		if isCredit {
			actions.Create("credit", name, value)
		}

		if isDebit {
			actions.Create("debit", name, value)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().BoolVarP(&isCredit, "credit", "c", false, "financas create -c")
	createCmd.Flags().BoolVarP(&isDebit, "debit", "d", false, "financas create -d")

	createCmd.Flags().StringVarP(&name, "name", "n", "", "financas create -n 'nome do registro'")
	createCmd.Flags().Float64VarP(&value, "value", "v", 0, "financas create -v 100.00")
}
