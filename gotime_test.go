package gotime

import (
	"testing"
	"fmt"
)

func TestGoTime_Now(t *testing.T) {
	tm := NewGoTime().Now()
	fmt.Println(tm)
}
func TestGoTime_Before(t *testing.T) {
	tm := NewGoTime().Before(3600)
	fmt.Println(tm)
}
func TestGoTime_Next(t *testing.T) {
	tm := NewGoTime().Next(3600)
	fmt.Println(tm)
}
func TestGoTime_GetYmd(t *testing.T) {
	tm := NewGoTime().GetYmd()
	fmt.Println(tm)
}
func TestGoTime_NowTime(t *testing.T) {
	tm := NewGoTime().NowTime()
	fmt.Println(tm)
}
func TestGoTime_RfcToUnix(t *testing.T) {
	tm := NewGoTime().RfcToUnix("2018-12-12T15:28:16+08:00")
	fmt.Println(tm)
}
func TestGoTime_GetRFC3339(t *testing.T) {
	tm := NewGoTime().GetRFC3339()
	fmt.Println(tm)
}
func TestGoTime_ToRFC3339(t *testing.T) {
	tm := NewGoTime().ToRFC3339(NewGoTime().Now())
	fmt.Println(tm)
}