package auth

type Auth struct {
	UserID     string `json:"userId"`     // 用户ID
	UserName   string `json:"userName"`   // 用户名称
	Source     string `json:"source"`     // 授权来源
	CreateTime string `json:"createTime"` // 授权时间
	Role       string `json:"role"`       // 角色
}
