package gotime

import (
	"time"
)

func NewBeijing() *GoTime {
	return &GoTime{
		Location: BeijingLocation,
	}
}

//获取当前时间 年-月-日 时:分:秒
func Now() string {
	return NewBeijing().Now()
}

//获取当前时间戳
func NowUnix() int64 {
	return NewBeijing().NowUnix()
}

//获取当前时间Time
func NowTime() time.Time {
	return NewBeijing().NowTime()
}

//获取年月日
func Ymd() string {
	return NewBeijing().Ymd()
}

//获取时分秒
func Hms() string {
	return NewBeijing().Hms()
}

//获取当天的开始时间, eg: 2018-01-01 00:00:00
func Start() string {
	return NewBeijing().Start()
}

//获取当天的结束时间, eg: 2018-01-01 23:59:59
func End() string {
	return NewBeijing().End()
}

//当前时间 减去 多少秒
func Before(beforeSecond int64) string {
	return NewBeijing().Before(beforeSecond)
}

//当前时间 加上 多少秒
func Next(beforeSecond int64) string {
	return NewBeijing().Next(beforeSecond)
}

//2006-01-02T15:04:05Z07:00 转 时间戳
func RfcToUnix(s string) int64 { //转化所需模板
	return NewBeijing().RfcToUnix(s)
}

//2006-01-02 15:04:05 转 时间戳
func ToUnix(s string) int64 {
	return NewBeijing().ToUnix(s)
}

//获取RFC3339格式
func GetRFC3339() string {
	return NewBeijing().GetRFC3339()
}

//转换成RFC3339格式
func ToRFC3339(s string) string {
	return NewBeijing().ToRFC3339(s)
}

//将RFC3339格式转成正常格式
func RFC3339To(s string) string {
	return NewBeijing().RFC3339To(s)
}

func GetFormat(format string) string {
	return NewBeijing().GetFormat(format)
}

// Format time.Time struct to string
// MM - month - 01
// M - month - 1, single bit
// DD - day - 02
// D - day 2
// YYYY - year - 2006
// YY - year - 06
// HH - 24 hours - 03
// H - 24 hours - 3
// hh - 12 hours - 03
// h - 12 hours - 3
// mm - minute - 04
// m - minute - 4
// ss - second - 05
// s - second = 5
func Format(t time.Time, format string) string {
	return NewBeijing().Format(t, format)
}

//将格式化为一天的零点
func ParseStart(s, fromFormat string) (result string, err error) {
	return NewBeijing().ParseStart(s, fromFormat)
}

//将格式化为一天最后一刻
func ParseEnd(s, fromFormat string) (result string, err error) {
	return NewBeijing().ParseEnd(s, fromFormat)
}
