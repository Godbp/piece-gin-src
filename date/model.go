package date

const (
	// 默认格式化
	DefaultFormDataTime = "2006-01-02 15:04:05"
	DefaultFormData     = "2006-01-02"
	DefaultFormTime     = "15:04:05"
)

const (
	// SecondsPerMinute 定义每分钟的秒数
	SecondsPerMinute = 60
	// SecondsPerHour 定义每小时的秒数
	SecondsPerHour = SecondsPerMinute * 60
	// SecondsPerDay 定义每天的秒数
	SecondsPerDay = SecondsPerHour * 24
)

// ResolveSeconds ResolveSeconds
type ResolveSeconds struct {
	Day    int64 `json:"day"`    // 天
	Hour   int64 `json:"hour"`   // 小时
	Minute int64 `json:"minute"` // 分钟
	Second int64 `json:"second"` // 秒
}
