package date

import "time"

// GetNowTimestamp 获取当前时间戳 秒s
func GetNowTimestamp() int64 {
	return time.Now().Unix()
}

// GetNowTimeMillisecond 获取当前时间戳 毫秒ms
func GetNowTimeMillisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

// GetNowTimeNanosecond 获取当前时间戳 纳秒
func GetNowTimeNanosecond() int64 {
	return time.Now().UnixNano()
}

// SecondConvertNanosecond 获取当前时间戳 纳秒 ns
func SecondConvertNanosecond(second int64) int64 {
	return second * 1e9
}

// SecondConvertMillisecond 获取当前时间戳 毫秒ms
func SecondConvertMillisecond(second int64) int64 {
	return second * 1e3
}

// GetNowTimestampByString 通过string转换成时间戳
func GetNowTimestampByString(dataStr string, format string) int64 {
	t, _ := time.Parse(format, dataStr)
	return t.Unix()
}
