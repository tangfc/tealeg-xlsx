package main

import (
	"go-test/excel/dbops"
	"github.com/labstack/gommon/log"
	"fmt"
)

func main() {

	downloadData, err := dbops.GetMemberInfo() // 获取需要导出数据
	if err != nil {
		log.Printf("%s", err)
		return
	}
	// 导出数据到Excel
	err = dbops.DownlaodExcel("./excel/data/斗罗大陆人员名单.xlsx", []string{"序号", "封号", "外号", "姓名", "性别", "签名"}, downloadData)
	if err != nil {
		log.Printf("%s", err)
		return
	}

	// 读取Excel数据
	readData, err := dbops.ReadExcel("./excel/data/斗罗大陆人员名单.xlsx")
	if err != nil {
		log.Printf("%s", err)
		return
	}
	fmt.Println(readData)
}
