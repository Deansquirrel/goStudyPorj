package main

import (
	"encoding/json"
	"fmt"
	"github.com/Deansquirrel/goStudyPorj/global"
)

func main() {
	//==========================================================
	//启动标识
	fmt.Println("程序启动")
	defer fmt.Println("程序退出")
	//==========================================================
	//加载配置文件
	sysConfig, err := global.GetConfig("config.toml")
	if err != nil {
		fmt.Println("配置文件加载错误。[", err, "]")
		return
	}
	//配置文件内容输出
	jsonConfig, err := json.Marshal(sysConfig)
	if err != nil {
		fmt.Println("配置文件转JSON时遇到错误。[", err, "]")
	} else {
		fmt.Println("配置文件", string(jsonConfig))
	}
	//==========================================================
	fmt.Println("==============================================")
	defer fmt.Println("==============================================")
	//==========================================================
	//调试代码

	//==========================================================
}
