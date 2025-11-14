package excel

import "github.com/tealeg/xlsx"

// reader reader of excel
type reader struct {
	// excel file name
	filename string

	// data
	items chan map[string]string
}

// NewExcelReader return ExcelReader with filename and chan length
func Reader(filename string, chanLen int) *reader {
	return &reader{
		filename: filename,
		items:    make(chan map[string]string, chanLen),
	}
}

// Read read excel sheet
func (r *reader) Read(sheetNos ...int /* sheet number*/) error {
	// open file
	xlFile, err := xlsx.OpenFile(r.filename)
	if err != nil {
		return err
	}

	sheetNo := 0
	if len(sheetNos) > 0 {
		sheetNo = sheetNos[0]
	}

	sheet := xlFile.Sheets[sheetNo]
	go r.load(sheet)

	return nil
}

func (r *reader) load(sheet *xlsx.Sheet) {
	// set first row as map filed
	var fieldArr []string

	for idx, row := range sheet.Rows {
		var arr []string
		// read data
		for _, cell := range row.Cells {
			arr = append(arr, cell.String())
		}

		if arr == nil { // filter empty row
			continue
		}

		// read filed
		if idx == 0 {
			fieldArr = arr
			continue
		}

		length := len(arr)
		item := make(map[string]string)
		for key, field := range fieldArr {
			if key >= length { // 处理out of range的问题
				continue
			}
			item[field] = arr[key]
		}

		r.items <- item
	}

	// close channel when read finished
	close(r.items)
}

// OutPut return read data
func (r *reader) OutPut() <-chan map[string]string {
	return r.items
}
