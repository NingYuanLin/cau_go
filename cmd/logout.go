package cmd

import (
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:     "logout",
	Aliases: []string{"o"},
	Short:   "注销用户",
	Run: func(cmd *cobra.Command, args []string) {
		Operation = OperationLogout
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
