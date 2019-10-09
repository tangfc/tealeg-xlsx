package dbops

import (
	"github.com/tealeg/xlsx"
	"go-test/excel/defs"
	"go-test/excel/comm"
	"github.com/labstack/gommon/log"
	"strconv"
)

var (
	file         *xlsx.File
	sheet        *xlsx.Sheet
	row          *xlsx.Row
	cell         *xlsx.Cell
	downloadData []*defs.MemberInfo
	err          error
)

// 下载Excel
func DownlaodExcel(filePath string, heads []string, data []*defs.MemberInfo) error {
	// 创建路径
	err = comm.CreateFilePath(filePath)
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}
	// 创建文件
	file = xlsx.NewFile()
	// 添加新工作表
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		return err
	}
	// 向工作表中添加新行
	row = sheet.AddRow()
	// 头部写入
	for _, head := range heads {
		cell = row.AddCell()
		cell.Value = head
	}
	// 设置单元格样式
	sheet.SetColWidth(5, 5, 60) // 设置单元格宽度 0-A 1-B 2-C

	// 主体写入数据
	for _, val := range data {
		row = sheet.AddRow()
		row.AddCell().Value = strconv.Itoa(val.Id) // 该行单元格写入数据
		row.AddCell().Value = val.Nickname
		row.AddCell().Value = val.Username
		row.AddCell().Value = val.Realname
		row.AddCell().Value = val.Sex
		row.AddCell().Value = val.Signature
	}
	// 在提供的路径中将文件保存到xlsx文件
	err = file.Save(filePath)
	if err != nil {
		return err
	}
	return nil
}
