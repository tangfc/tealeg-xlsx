package dbops

import (
	"github.com/tealeg/xlsx"
)

// 读取EXCEL
func ReadExcel(filePath string) (data map[string][][]string, err error) {
	data = make(map[string][][]string)
	xlFile, err := xlsx.OpenFile(filePath) // 打开文件
	if err != nil {
		return
	}
	for _, sheet := range xlFile.Sheets { // 循环打开工作簿
		var sheetInfo [][]string
		for _, row := range sheet.Rows { // 循环读取行
			var info []string
			for _, cell := range row.Cells { // 循环读取单元格
				info = append(info, cell.String())
			}
			sheetInfo = append(sheetInfo, info)
		}
		data[sheet.Name] = sheetInfo
	}
	return
}
