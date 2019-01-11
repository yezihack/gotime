package gotime

import (
	"fmt"
	"testing"
	"time"
)

func TestGoTime_Now(t *testing.T) {
	tm := Now()
	fmt.Println(tm)
}
func TestGoTime_Before(t *testing.T) {
	tm := Before(3600)
	fmt.Println(tm)
}
func TestGoTime_Next(t *testing.T) {
	tm := Next(3600)
	fmt.Println(tm)
}
func TestGoTime_GetYmd(t *testing.T) {
	tm := Ymd()
	fmt.Println(tm)
}
func TestGoTime_NowTime(t *testing.T) {
	tm := NowTime()
	fmt.Println(tm)
}
func TestGoTime_RfcToUnix(t *testing.T) {
	tm := RfcToUnix("2018-12-12T15:28:16+08:00")
	fmt.Println(tm)
}
func TestGoTime_GetRFC3339(t *testing.T) {
	tm := GetRFC3339()
	fmt.Println(tm)
}
func TestGoTime_ToRFC3339(t *testing.T) {
	tm := ToRFC3339(Now())
	fmt.Println(tm)
}
func TestGoTime_GetFormat(t *testing.T) {
	tm := GetFormat("YYYYMM")
	fmt.Println(tm)
}
//自定义时间
func TestSelf(t *testing.T) {
	loc, err := time.LoadLocation("America/Cordoba")
	loc, err = time.LoadLocation("Asia/Chongqing")
	if err != nil {
		panic(err)
	}
	g := GoTime{
		Location:loc,
	}
	fmt.Println(g.Now())
	fmt.Println(g.GetRFC3339())
}