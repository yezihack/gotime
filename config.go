package gotime

import "time"

//常量
const (
	RFC3339 = "2006-01-02T15:04:05+08:00"
	TT      = "2006-01-02 15:04:05"
	YMD     = "2006-01-02"
	HMS     = "15:04:05"
)

//北京东八时区
var BeijingLocation = time.FixedZone("Asia/Shanghai", 8*60*60)
