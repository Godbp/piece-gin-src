package auth

const (
	LoginAuth  = "loginAuth"  // 用户登录授权
	AccessAuth = "accessAuth" // 访问授权
)

// 基础用户信息
type BaseAuth struct {
	UserID     string `json:"userId"`     // 用户ID
	UserName   string `json:"userName"`   // 用户名称
	Source     string `json:"source"`     // 授权来源
	CreateTime string `json:"createTime"` // 授权时间
	Role       string `json:"role"`       // 角色
	AuthType   string `json:"authType"`   // 授权类型
}

// 用户授权信息补全
type Auth struct {
	BaseAuth
}
