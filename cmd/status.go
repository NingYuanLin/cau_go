package cmd

import "github.com/spf13/cobra"

var statusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"s"},
	Short:   "查询登录状态",
	Run: func(cmd *cobra.Command, args []string) {
		Operation = OperationStatus
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
