package pkgs

import "time"

// String2Time 时间字符串转 time
func String2Time(data string) time.Time {
	loc, _ := time.LoadLocation("Local")
	checkdate, err := time.ParseInLocation("2006-01-02 15:04:05", data, loc)
	if err != nil {
		return time.Time{}
	}
	return checkdate
}

// Time2String 时间转字符串
func Time2String(data time.Time) string {
	timeString := data.Format("2006-01-02 15:04:05")
	return timeString
}
