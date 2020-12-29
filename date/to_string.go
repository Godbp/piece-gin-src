package date

import (
	"math"
	"strconv"
	"strings"
	"time"
)

// GetNowDateTimeFormatByFormat 返回指定格式的日期字符串
func GetNowDateTimeFormatByFormat(format string) string {
	t := time.Now()
	return t.Format(format)
}

// UtcToLocal 返回utc当前时间
func UtcToLocal(utcTime string) string {
	layout := "2006-01-02T15:04:05Z"
	utc, _ := time.LoadLocation("UTC")
	newTime, _ := time.ParseInLocation(layout, utcTime, utc)

	return newTime.Local().Format(DefaultFormDataTime)
}

// GetNowTime 获取当前时间
func GetNowTime(t time.Time) (now string) {
	hour := strconv.Itoa(t.Hour())
	min := strconv.Itoa(t.Minute())
	return hour + ":" + min
}

// GetNowDateFormat 获取当前格式化日期
func GetNowDateFormat() string {
	t := time.Now()
	return t.Format("2006-01-02")
}

// GetNowDateTimeFormat 获取当前格式化的时间字符串
func GetNowDateTimeFormat() string {
	t := time.Now()
	return t.Format(DefaultFormDataTime)
}

// GetNowDateTimeFormatCustom 传入指定格式格式化当前时间
func GetNowDateTimeFormatCustom(format string) string {
	t := time.Now()
	return t.Format(format)
}

// GetTimeFormat 根据时间戳转换格式化日期时间
func GetTimeFormat(ts int64) string {
	t := time.Unix(ts, 0)
	return t.Format(DefaultFormDataTime)
}

// GetDateFormat 根据时间类型转换格式化日期
func GetDateFormat(t time.Time) string {
	return t.Format(DefaultFormData)
}

// GetDateTimeFormat 获取当前格式化的时间字符串
func GetDateTimeFormat(t time.Time) string {
	return t.Format(DefaultFormDataTime)
}

// GetYearMonthFormat 获取年月
func GetYearMonthFormat(t time.Time) string {
	str := t.Format(DefaultFormData)
	strArr := strings.Split(str, "-")
	return strArr[0] + "-" + strArr[1]
}

// GetYear 获取年月
func GetYear(t time.Time) string {
	str := t.Format(DefaultFormData)
	strArr := strings.Split(str, "-")
	return strArr[0]
}

// GetDateTimeBeforeOrAfter 获取几天前 年月日 时分秒格式
func GetDateTimeBeforeOrAfter(t time.Time, days int) string {
	t = t.AddDate(0, 0, days)
	return t.Format(DefaultFormDataTime)
}

// GetDaysBeforeOrAfter 获取几天前或者几天后，年月日格式
func GetDaysBeforeOrAfter(t time.Time, days int) string {
	t = t.AddDate(0, 0, days)
	return t.Format("2006-01-02")
}

// GetMonthsBeforeOrAfter 获得几个月前或者几个月后的年月格式
func GetMonthsBeforeOrAfter(t time.Time, months int) string {
	t = t.AddDate(0, months, 0)
	return t.Format("2006-01")
}

// GetWeeksBeforeOrAfter 获得几周前或者几周后的年月日格式
func GetWeeksBeforeOrAfter(t time.Time, weeks int) string {
	t = t.AddDate(0, 0, weeks*7)
	offset := int(t.Weekday()-7) / -1
	day := time.Date(t.Year(), t.Month(), t.Day()+offset, 0, 0, 0, 0, t.Location())
	return day.Format("2006-01-02")
}

// GetStartOfWeekFormat 获取本周的开始时间并且格式化
func GetStartOfWeekFormat(t time.Time) string {
	offset := int(t.Weekday()+6) % 7
	startTime := time.Date(t.Year(), t.Month(), t.Day()-offset, 0, 0, 0, 0, t.Location())
	return startTime.Format(DefaultFormDataTime)
}

// GetStartOfWeekFormatDate 获取本周的开始时间并且格式化 年-月-日
func GetStartOfWeekFormatDate(t time.Time) string {
	offset := int(t.Weekday()+6) % 7
	startTime := time.Date(t.Year(), t.Month(), t.Day()-offset, 0, 0, 0, 0, t.Location())
	return startTime.Format("2006-01-02")
}

// GetEndOfWeekFormat 获取本周的结束时间并且格式化
func GetEndOfWeekFormat(t time.Time) string {
	weekend := t.Weekday()
	if weekend == 0 {
		weekend = 7
	}
	offset := int(weekend-7) / -1
	startTime := time.Date(t.Year(), t.Month(), t.Day()+offset, 0, 0, 0, 0, t.Location())
	return startTime.Format(DefaultFormDataTime)
}

// GetEndOfWeekFormatDate 获取本周的结束时间并且格式化 年-月-日
func GetEndOfWeekFormatDate(t time.Time) string {
	weekend := t.Weekday()
	if weekend == 0 {
		weekend = 7
	}
	offset := int(weekend-7) / -1
	startTime := time.Date(t.Year(), t.Month(), t.Day()+offset, 0, 0, 0, 0, t.Location())
	return startTime.Format("2006-01-02")
}

// GetStartOfDayFormat 获取一天的开始时间并且格式化
func GetStartOfDayFormat(t time.Time) string {
	startTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return startTime.Format(DefaultFormDataTime)
}

// GetDaysBeforeOrAfterFormat 获取几天前的零点时间
func GetDaysBeforeOrAfterFormat(t time.Time, days int) string {
	t = t.AddDate(0, 0, days)
	startTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return startTime.Format(DefaultFormDataTime)
}

// GetDaysBeforeMouthFormat 获取上个月的零点时间
func GetDaysBeforeMouthFormat(t time.Time, days int) string {
	t = t.AddDate(0, 0, days)
	startTime := time.Date(t.Year(), t.Month()-1, t.Day(), 0, 0, 0, 0, t.Location())
	return startTime.Format(DefaultFormDataTime)
}

// GetWeekDay 获取星期几
func GetWeekDay(t time.Time) string {
	weekDay := t.Weekday()
	res := ""
	switch weekDay {
	case time.Sunday:
		res = "7"
	case time.Monday:
		res = "1"
	case time.Tuesday:
		res = "2"
	case time.Wednesday:
		res = "3"
	case time.Thursday:
		res = "4"
	case time.Friday:
		res = "5"
	case time.Saturday:
		res = "6"
	}
	return res
}

// GetTimeBySecond 根据秒转换时间
func GetTimeBySecond(second float64) string {
	hour := int(math.Floor(second / 3600))
	minute := int(math.Floor(math.Mod(second, 3600) / 60))
	sec := int(math.Mod(second, 60))

	hourStr := strconv.Itoa(hour)
	minuteStr := strconv.Itoa(minute)
	secStr := strconv.Itoa(sec)
	if minute < 10 {
		minuteStr = "0" + minuteStr
	}
	if sec < 10 {
		secStr = "0" + secStr
	}
	return hourStr + ":" + minuteStr + ":" + secStr
}

// GetShowTimeBySecond 根据秒转换时间-时分秒
func GetShowTimeBySecond(second float64) string {
	hour := int(math.Floor(second / 3600))
	minute := int(math.Floor(math.Mod(second, 3600) / 60))
	sec := int(math.Mod(second, 60))

	hourStr := strconv.Itoa(hour)
	minuteStr := strconv.Itoa(minute)
	secStr := strconv.Itoa(sec)
	// if minute < 10 {
	// 	minuteStr = "0" + minuteStr
	// }
	// if sec < 10 {
	// 	secStr = "0" + secStr
	// }
	return hourStr + "时" + minuteStr + "分钟" + secStr + "秒"
}

// GetSecondByTime 根据时间转换秒
func GetSecondByTime(time string) string {
	if time == "" {
		return "0"
	}
	times := strings.Split(time, ":")
	if len(times) == 3 {
		hour, _ := strconv.Atoi(times[0])
		minute, _ := strconv.Atoi(times[1])
		second, _ := strconv.Atoi(times[2])
		return strconv.Itoa(hour*3600 + minute*60 + second)
	}
	return "0"
}

// ToDateTime 转换时间
func ToDateTime(date []string) []string {
	dateTime := make([]string, 0, 2)
	if len(date) == 2 {
		startTime := date[0] + " 00:00:00"
		endTime := date[1] + " 23:59:59"
		dateTime = append(dateTime, startTime)
		dateTime = append(dateTime, endTime)
		return dateTime
	}
	return date
}

// GetShowTime 长短时间
func GetShowTime() (string, string) {
	nowCom := time.Now()
	nowLong := nowCom.Format(DefaultFormDataTime)
	nowShort := nowCom.Format("2006-01-02 15:04")
	return nowLong, nowShort
}

// TimeStrToUnix 字符串转时间戳
func TimeStrToUnix(timeStr string) (int64, error) {
	timeLayout := "2006-01-02 15:04"       // 转化所需模板
	loc, err := time.LoadLocation("Local") // 重要：获取时区
	if err != nil {
		return 0, err
	}
	theTime, err := time.ParseInLocation(timeLayout, timeStr, loc) // 使用模板在对应时区转化为time.time类型
	if err != nil {
		return 0, err
	}
	sr := theTime.Unix()
	return sr, nil
}

// TimeStrBeforeAfterDay 几天前/几天后的时间
func TimeStrBeforeAfterDay(timeStr string, days int) (string, error) {
	timeLayout := "2006-01-02 15:04"       // 转化所需模板
	loc, err := time.LoadLocation("Local") // 重要：获取时区
	if err != nil {
		return "0", err
	}
	theTime, err := time.ParseInLocation(timeLayout, timeStr, loc) // 使用模板在对应时区转化为time.time类型
	if err != nil {
		return "", err
	}
	return theTime.AddDate(0, 0, days).Format(timeLayout), nil
}

// TimeStrBeforeAfterHour 几小时前/后的时间
func TimeStrBeforeAfterHour(timeStr string, hours int, minutes int) (string, error) {
	timeLayout := "2006-01-02 15:04"       // 转化所需模板
	loc, err := time.LoadLocation("Local") // 重要：获取时区
	if err != nil {
		return "0", err
	}
	theTime, err := time.ParseInLocation(timeLayout, timeStr, loc) // 使用模板在对应时区转化为time.time类型
	if err != nil {
		return "", err
	}
	year, month, day := theTime.Date()
	hour, min, sec := theTime.Clock()
	theTime = time.Date(year, month, day, hour+hours, min+minutes, sec, 0, theTime.Location())
	return theTime.Format(timeLayout), nil
}

// TimeStrToUnixByFormat 字符串转时间戳
func TimeStrToUnixByFormat(timeStr, format string) (int64, error) {
	loc, err := time.LoadLocation("Local") // 重要：获取时区
	if err != nil {
		return 0, err
	}

	// 使用模板在对应时区转化为time.time类型
	theTime, err := time.ParseInLocation(format, timeStr, loc)
	if err != nil {
		return 0, err
	}

	return theTime.Unix(), nil
}

// RemainingSecondCurrentDay 返回今天剩余秒数
func RemainingSecondCurrentDay() float64 {
	now := time.Now()
	t1 := now.AddDate(0, 0, 1)
	t2 := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
	return t2.Sub(now).Seconds()
}

// ResolveTime 解析秒为*天*小时*分钟*秒
func ResolveTime(seconds int64) *ResolveSeconds {
	resolveSeconds := &ResolveSeconds{}
	resolveSeconds.Day = seconds / SecondsPerDay

	if seconds%SecondsPerDay != 0 {
		resolveSeconds.Hour = seconds % SecondsPerDay / SecondsPerHour
	}

	if seconds%SecondsPerDay%SecondsPerHour != 0 {
		resolveSeconds.Minute = seconds % SecondsPerDay % SecondsPerHour / SecondsPerMinute
	}

	if seconds%SecondsPerDay%SecondsPerHour%SecondsPerMinute != 0 {
		resolveSeconds.Second = seconds % SecondsPerDay % SecondsPerHour % SecondsPerMinute % 60
	}
	return resolveSeconds
}

// GetWeekDayInt 获取星期几
func GetWeekDayInt(t time.Time) int {
	weekDay := t.Weekday()
	var res int
	switch weekDay {
	case time.Sunday:
		res = 0
	case time.Monday:
		res = 1
	case time.Tuesday:
		res = 2
	case time.Wednesday:
		res = 3
	case time.Thursday:
		res = 4
	case time.Friday:
		res = 5
	case time.Saturday:
		res = 6
	}
	return res
}

// GetTimeWithLayout 时间戳转化
func GetTimeWithLayout(ts int64, layout string) string {
	t := time.Unix(ts, 0).Local()
	return t.Format(layout)
}
