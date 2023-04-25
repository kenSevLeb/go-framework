package date

import (
	"fmt"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"
)

// GetLocalTime 返回东8时区的时间
func GetLocalTime() time.Time {
	return time.Now().In(GetLocalTimeZone())
}

// GetLocalTimeZone
func GetLocalTimeZone() *time.Location {
	return time.FixedZone("CST", 8*3600) // UTC+8
}

// GetLocalMicroTimeStampStr 返回ms
func GetLocalMicroTimeStampStr() string {
	return fmt.Sprintf("%.6f", float64(GetLocalTime().UnixNano())/1e9)
}

// 解析时间
func ParseTime(timeStr string) time.Time {
	tmp, _ := time.ParseInLocation(TimeFormat, timeStr, GetLocalTimeZone())
	return tmp
}

// 解析日期
func ParseDate(dateStr string) time.Time {
	tmp, _ := time.ParseInLocation(DateFormat, dateStr, GetLocalTimeZone())
	return tmp
}

// 时间戳
func Timestamp() int64 {
	return GetLocalTime().Unix()
}

// 获取时间
func GetTimeStr() string {
	return GetLocalTime().Format(TimeFormat)
}

// 获取日期
func GetDateStr() string {
	return GetLocalTime().Format(DateFormat)
}

// 时间戳转时间
func FormatTime(timestamp int64) string {
	return time.Unix(timestamp, 0).In(GetLocalTimeZone()).Format(TimeFormat)
}

// 时间戳转日期
func FormatDate(timestamp int64) string {
	return time.Unix(timestamp, 0).In(GetLocalTimeZone()).Format(DateFormat)
}
