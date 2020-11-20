package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-----------------")
	fmt.Println(time.Now().Hour())
	var now = time.Now()
	fmt.Println(now.UnixNano())
	var now1 = time.Since(now)
	fmt.Println(time.Now().UnixNano())
	fmt.Println(now1.Nanoseconds())
	GetLocalSomeDayZeroTime(0, 0, 0)
	fmt.Println("-----------------")
	getLocalTodayZeroTime()
}

func getLocalTodayZeroTime() (zeroTime time.Time) {
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println(timeStr)
	zeroTime, _ = time.ParseInLocation("2006-01-02", timeStr, time.FixedZone("UTC", 8*3600))
	tt, _ := time.Parse("2006-01-02", timeStr)
	fmt.Println(zeroTime, zeroTime.Unix())
	fmt.Println(tt, tt.Unix())
	return
}

// GetLocalSomeDayZeroTime
func GetLocalSomeDayZeroTime(year int, month int, day int) (timeUnix int64) {

	timeUnix = getLocalTodayZeroTime().AddDate(year, month, day).Unix()
	ts := getLocalTodayZeroTime().Format("2006_01_02")
	fmt.Println(ts)
	return
}

func format() {
	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	t = t.AddDate(0, 0, -1)
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}
