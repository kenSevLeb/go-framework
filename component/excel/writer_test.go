package excel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const filename = "/home/hongker/Downloads/target01.xlsx"

func TestWriter(t *testing.T) {
	w := Writer(filename)
	err := w.Write([]string{"标题", "内容"}, [][]string{
		{"hello", "world"},
		{"hello", ""},
		{"", "world"},
	})

	assert.Nil(t, err)
}

func TestWriter_WriteWithSheet(t *testing.T) {
	w := Writer(filename)
	err := w.WriteWithSheet([]string{"标题", "内容"}, [][]string{
		{"hello", "world"},
		{"hello", ""},
		{"", "world"},
	}, "Sheet2")

	assert.Nil(t, err)
}
