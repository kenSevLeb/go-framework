package date

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetLocalTime(t *testing.T) {
	fmt.Println(GetLocalTime().Unix())
}

func TestGetLocalMicroTimeStampStr(t *testing.T) {
	fmt.Println(GetLocalMicroTimeStampStr())
}

func TestMonth(t *testing.T) {
	fmt.Println(int(GetLocalTime().Month()))
}

func TestParseDate(t *testing.T) {
	fmt.Println(ParseDate("2020-09-10").String())
}

func TestParseTime(t *testing.T) {
	fmt.Println(ParseTime("2020-09-10 11:22:33").String())
}

func TestTimestamp(t *testing.T) {
	now := time.Now().Unix()
	assert.Equal(t, now, Timestamp())
}

func TestGetTimeStr(t *testing.T) {
	fmt.Println(GetTimeStr())
}

func TestGetDateStr(t *testing.T) {
	fmt.Println(GetDateStr())
}

func TestFormatTime(t *testing.T) {
	fmt.Println(FormatTime(Timestamp()))
}

func TestFormatDate(t *testing.T) {
	fmt.Println(FormatDate(Timestamp()))
}
