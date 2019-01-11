package gotime

import (
	"time"
)

type GoTime struct {
	Location *time.Location
}

//获取当前时间 年-月-日 时:分:秒
func (g *GoTime) Now() string {
	return g.NowTime().Format(TT)
}

//获取当前时间戳
func (g *GoTime) NowUnix() int64 {
	return g.NowTime().Unix()
}

//获取当前时间Time
func (g *GoTime) NowTime() time.Time {
	return time.Now().In(g.Location)
}

//获取年月日
func (g *GoTime) Ymd() string {
	return g.NowTime().Format(YMD)
}

//获取时分秒
func (g *GoTime) Hms() string {
	return g.NowTime().Format(HMS)
}

//获取当天的开始时间, eg: 2018-01-01 00:00:00
func (g *GoTime) Start() string {
	now := g.NowTime()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, g.Location)
	return tm.Format(TT)
}

//获取当天的结束时间, eg: 2018-01-01 23:59:59
func (g *GoTime) End() string {
	now := g.NowTime()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, g.Location)
	return tm.Format(TT)
}

//当前时间 减去 多少秒
func (g *GoTime) Before(beforeSecond int64) string {
	return time.Unix(g.NowUnix()-beforeSecond, 0).Format(TT)
}

//当前时间 加上 多少秒
func (g *GoTime) Next(beforeSecond int64) string {
	return time.Unix(g.NowUnix()+beforeSecond, 0).Format(TT)
}

//2006-01-02T15:04:05Z07:00 转 时间戳
func (g *GoTime) RfcToUnix(s string) int64 { //转化所需模板
	tm, err := time.ParseInLocation(time.RFC3339, s, g.Location) //使用模板在对应时区转化为time.time类型
	if err != nil {
		return int64(0)
	}
	return tm.Unix()
}

//2006-01-02 15:04:05 转 时间戳
func (g *GoTime) ToUnix(s string) int64 {
	t, _ := time.ParseInLocation(TT, s, g.Location)
	return t.Unix()
}

//获取RFC3339格式
func (g *GoTime) GetRFC3339() string {
	return g.NowTime().Format(time.RFC3339)
}

//转换成RFC3339格式
func (g *GoTime) ToRFC3339(s string) string {
	tm, err := time.ParseInLocation(TT, s, g.Location)
	if err != nil {
		return ""
	}
	return tm.Format(time.RFC3339)
}

//将RFC3339格式转成正常格式
func (g *GoTime) RFC3339To(s string) string {
	t, err := time.ParseInLocation(RFC3339, s, g.Location)
	if err != nil {
		return "0000-00-00 00:00:00"
	}
	return t.Format(TT)
}

//获取格式化的数据
func (g *GoTime) GetFormat(format string) string {
	return g.Format(g.NowTime(), format)
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
func (g *GoTime) Format(t time.Time, format string) string {
	return Format(t, format)
}

//将格式化为一天的零点
func (g *GoTime) ParseStart(s, fromFormat string) (result string, err error) {
	tt, err := time.Parse(fromFormat, s)
	if err != nil {
		return
	}
	tt = time.Date(tt.Year(), tt.Month(), tt.Day(), 0, 0, 0, 0, g.Location)
	return tt.Format(TT), nil
}

//格式化为一天的最后一刻
func (g *GoTime) ParseEnd(s, fromFormat string) (result string, err error) {
	tt, err := time.Parse(fromFormat, s)
	if err != nil {
		return
	}
	tt = time.Date(tt.Year(), tt.Month(), tt.Day(), 23, 59, 59, 0, g.Location)
	return tt.Format(TT), nil
}
