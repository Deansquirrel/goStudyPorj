package main

import (
	"encoding/json"
	"fmt"
	"github.com/Deansquirrel/go-tool"
	"github.com/Deansquirrel/goStudyPorj/global"
	"github.com/Deansquirrel/goStudyPorj/lib"
	"time"
)

func main() {
	//==========================================================
	//启动标识
	fmt.Println("程序启动")
	defer fmt.Println("程序退出")
	//==========================================================
	fmt.Println("==============================================")
	defer fmt.Println("==============================================")
	//==========================================================
	//调试代码
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
	tickets, err := lib.NewMyGoTickets(5)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for i := 0; i < 10; i++ {
			tickets.Take()
			fmt.Println(go_tool.GetDateTimeStr(time.Now()), "[tke]")
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 10)
	for i := 0; i < 10; i++ {
		tickets.Return()
		fmt.Println(go_tool.GetDateTimeStr(time.Now()), "[return]")
		time.Sleep(time.Second * 2)
	}
	//==========================================================
}
