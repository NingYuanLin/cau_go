package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var loginCmd = &cobra.Command{
	Use:     "login",
	Aliases: []string{"i"},
	Short:   "登录。当不指定-u和-p时，会尝试读取\"$HOME/.cau_go.yaml\"的配置文件",
	Args:    cobra.OnlyValidArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		// 将initConfig()，也就是读取配置文件放到这里面，就可以当执行login command的时候才调用了
		// 如果我们在init()中使用cobra.OnInitialize(initConfig),无论command是什么，在调用rootCmd.Execute()的时候，都会调用initConfig
		// 这个时候 其实argument已经解析好了，但是不用担心在这里面执行initConfig会覆盖掉原来的参数
		// 这是因为：当我们执行viper.Get的时候，
		// Viper will then check in the following order:
		// flag, env, config file, key/value store.
		// 所以，当Flag存在的时候，会优先读取flag
		initConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("args:", args)
		//fmt.Println("username:", viper.GetString("username"))
		//viper.Set("Operation", OperationLogin)
		Operation = OperationLogin
	},
}

var cfgFile string

func init() {
	// 放到PreRun()里执行了
	//cobra.OnInitialize(initConfig)

	loginCmd.PersistentFlags().StringP("username", "u", "", "用户名")
	loginCmd.PersistentFlags().StringP("password", "p", "", "密码")
	loginCmd.MarkFlagsRequiredTogether("username", "password")
	loginCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.cau_go.yaml)")
	loginCmd.MarkFlagsMutuallyExclusive("config", "username")
	loginCmd.MarkFlagsMutuallyExclusive("config", "password")

	viper.BindPFlag("username", loginCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", loginCmd.PersistentFlags().Lookup("password"))

	rootCmd.AddCommand(loginCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	//fmt.Println("initConfig")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		//home, err := os.UserHomeDir()
		// 使用go-homedir库可以支持交叉编译
		home, err := homedir.Dir()
		//print(home)
		cobra.CheckErr(err)

		// Search config in home directory with name ".cau_go" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cau_go")
	}

	//viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		//fmt.Println("config not find")
	}
}
