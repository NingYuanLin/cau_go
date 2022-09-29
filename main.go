/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/NingYuanLin/cau_go/cmd"
	"github.com/NingYuanLin/cau_go/core/config"
	"github.com/NingYuanLin/cau_go/core/login"
	"github.com/NingYuanLin/cau_go/core/logout"
	"github.com/NingYuanLin/cau_go/core/status"
)

func main() {
	cmd.Execute()
	// 默认返回0

	switch cmd.Operation {
	case cmd.OperationLogin:
		err := login.Run()
		if err != nil {
			fmt.Println(err.Error())
			//fmt.Println("↑↑↑ 发生错误 ↑↑↑")
			return
		}
	case cmd.OperationConfig:
		err := config.Run()
		if err != nil {
			fmt.Println(err.Error())
			//fmt.Println("↑↑↑ 发生错误 ↑↑↑")
			return
		}
	case cmd.OperationLogout:
		err := logout.Run()
		if err != nil {
			fmt.Println(err.Error())
			//fmt.Println("↑↑↑ 发生错误 ↑↑↑")
			return
		}
	case cmd.OperationStatus:
		err := status.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	default:
		//fmt.Println("请输出正确的指令！")
	}

}
