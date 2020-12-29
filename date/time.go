package date

import "time"

// GetDataSubValue 计算两个日期的差值
func GetDataSubValue(newData string) (day, hour, minute, second float64, err error) {
	nowData := GetNowDateTimeFormat()
	if len(newData) == 10 {
		newData += " 00:00:00"
		nowData = GetNowDateFormat() + " 00:00:00"
	}

	loc, err := time.LoadLocation("Local")
	if err != nil {
		return
	}
	newTime, err := time.ParseInLocation(DefaultFormDataTime, newData, loc)
	if err != nil {
		return
	}
	nowTime, err := time.ParseInLocation(DefaultFormDataTime, nowData, loc)
	if err != nil {
		return
	}

	dtn := newTime.Sub(nowTime)
	hour = dtn.Hours()
	minute = dtn.Minutes()
	second = dtn.Seconds()
	day = hour / 24
	return
}

// GetTimeItem 获取时间的时分秒
func GetTimeItem(t *time.Time) (hour int, min int, sec int) {
	hour = t.Hour()
	min = t.Minute()
	sec = t.Second()
	return
}

// GetDateItem 获取日期的年月日
func GetDateItem(t *time.Time) (year int, month int, day int) {
	year = t.Year()
	month = int(t.Month())
	day = t.Day()
	return
}

// GetTimeFromDefaultString 根据日期字符串转换时间
func GetTimeFromDefaultString(str string) time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, str)
	return t
}

// GetTimeFromDefaultStringFull 根据日期字符串转换时间
func GetTimeFromDefaultStringFull(str string) time.Time {
	layout := "2006-01-02 15:04:05"

	t, _ := time.Parse(layout, str)
	return t
}

// IsLeapYear 判断是否为闰年
func IsLeapYear(year int) bool { // y == 2000, 2004
	// 判断是否为闰年
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}

	return false
}

// GetFirstDateOfMonth 获取一个月的的第一天
func GetFirstDateOfMonth(t time.Time) time.Time {
	t = t.AddDate(0, 0, -t.Day()+1)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// GetLastDateOfMonth 获取一个月的最后一天
func GetLastDateOfMonth(t time.Time) time.Time {
	return GetFirstDateOfMonth(t).AddDate(0, 1, -1)
}

// GetStartOfDay 获取一天的开始时间
func GetStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// GetMonthDays 获取某月有多少天
func GetMonthDays(year int, month int) int {
	days := 0
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30
		} else {
			days = 31
		}
	} else {
		if IsLeapYear(year) {
			days = 29
		} else {
			days = 28
		}
	}
	return days
}
