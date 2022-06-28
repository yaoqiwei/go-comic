package time

import "time"

// 获取当前的时间 - 字符串
func GetCurrentDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetCurrentDay() string {
	return time.Now().Format("2006-01-02")
}

func Format(i int64) string {
	return time.Unix(i, 0).Format("2006-01-02 15:04:05")
}

// 获取当前的时间 - Unix时间戳
func GetCurrentUnix() int64 {
	return time.Now().Unix()
}

// 获取当前的时间 - 毫秒级时间戳
func GetCurrentMilliUnix() int64 {
	return time.Now().UnixNano() / 1000000
}

// 获取当前的时间 - 纳秒级时间戳
func GetCurrentNanoUnix() int64 {
	return time.Now().UnixNano()
}

func GetWeekStartUnix() int64 {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset).Unix()
}

func GetWeekEndUnix() int64 {
	return GetWeekStartUnix() + 3600*24*7
}

func GetMonthStartUnix() int64 {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	return time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation).Unix()
}

func GetMonthEndUnix() int64 {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	return time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation).AddDate(0, 1, 0).Unix()
}

func GetDayStartUnix() int64 {
	now := time.Now()
	currentYear, currentMonth, currentDay := now.Date()
	currentLocation := now.Location()
	return time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation).Unix()
}

func GetDayEndUnix() int64 {
	return GetDayStartUnix() + 3600*24
}

func GetTimeFromUinx(second int64) time.Time {
	return time.Unix(second, 0)
}
