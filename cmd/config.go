/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "配置文件。-c创建配置文件",
	PreRun: func(cmd *cobra.Command, args []string) {
		initConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		//getBool, err := cmd.Flags().GetBool("create")
		//viper.Set("Operation", OperationConfig)
		Operation = OperationConfig
		if doConfigCreateBool == true {
			DoConfigOperationFlag = DoConfigCreateFlag
		} else {
			// 默认操作
			DoConfigOperationFlag = DoConfigNoneFlag
		}
	},
}
var doConfigCreateBool bool

var DoConfigOperationFlag int

const (
	DoConfigNoneFlag int = iota
	DoConfigCreateFlag
)

func init() {
	//cobra.OnInitialize(initConfig)
	configCmd.Flags().BoolVarP(&doConfigCreateBool, "create", "c", false, "创建配置文件")
	rootCmd.AddCommand(configCmd)
}
