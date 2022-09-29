package config

import (
	"bufio"
	"fmt"
	"github.com/NingYuanLin/cau_go/cmd"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func Run() error {
	switch cmd.DoConfigOperationFlag {
	case cmd.DoConfigCreateFlag:
		// 创建配置文件
		err := createConfig()
		if err != nil {
			return err
		}
	default:
		fmt.Println("请输入正确的flag，在命令后加--help获取帮助")
	}
	return nil
}

func createConfig() error {
	// 判断配置文件是否存在
	err := viper.ReadInConfig()
	configFileExisted := false
	if err == nil {
		// 配置文件存在
		// TODO:viper无法获取当前设置的配置文件路径
		fmt.Println("注意！配置文件已经存在,继续将会覆盖原有配置文件")
		configFileExisted = true
		//return err
	}
	fmt.Print("请输入您校园网的登录用户名：")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	username = strings.TrimSpace(username)

	fmt.Print("请输入您校园网的登录密码：")
	reader = bufio.NewReader(os.Stdin)
	password, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	password = strings.TrimSpace(password)

	viper.Set("username", username)
	viper.Set("password", password)
	if configFileExisted == true {
		err = viper.WriteConfig()
	} else {
		err = viper.SafeWriteConfig()
	}
	if err != nil {
		return err
	}
	return nil
}
