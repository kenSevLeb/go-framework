package excel

import (
	"fmt"
	"kenSevLeb/go-framework/util/file"
	"kenSevLeb/go-framework/util/strings"
	"github.com/tealeg/xlsx"
	"path/filepath"
)

const (
	defaultSheet = "Sheet1"
)

type writer struct {
	// 文件
	filename string
}

// Writer excel写入器
func Writer(filename string) *writer {
	return &writer{
		filename: filename,
	}
}

// Write 写入excel文件，默认写入sheet名称为Sheet1的表格
// headers : 头部，第一行，可以为nil
// items: 数据项
func (w *writer) Write(headers []string, items [][]string) error {
	return w.WriteWithSheet(headers, items, defaultSheet)
}

// WriteWithSheet 写入excel文件
// headers : 头部，第一行，可以为nil
// items: 数据项
// sheetName: sheet的名称
func (w *writer) WriteWithSheet(headers []string, items [][]string, sheetName string) error {
	if err := file.Mkdir(filepath.Dir(w.filename), true); err != nil {
		return fmt.Errorf("create directory:%s", err.Error())
	}

	var err error
	f := xlsx.NewFile()
	if file.Exist(w.filename) {

		f, err = xlsx.OpenFile(w.filename)
		if err != nil {
			return err
		}
	}
	sheet := f.Sheet[sheetName] // 检查是否已存在
	if sheet == nil {
		sheet, err = f.AddSheet(strings.Default(sheetName, defaultSheet))
		if err != nil {
			return err
		}
	}

	header := sheet.AddRow()
	for _, item := range headers {
		header.AddCell().Value = item
	}

	for _, item := range items {
		row := sheet.AddRow()
		for _, v := range item {
			row.AddCell().Value = v
		}
	}

	if err := f.Save(w.filename); err != nil {
		return err
	}

	return nil
}
