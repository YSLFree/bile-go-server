package code

const (
	StatusFailed          = 0 //失败
	StatusSuccess         = 1 //成功
	StatusUnkonwnErr      = 2 //未知错误
	StatusTokenInvalid    = 3 //token失效，身份无效
	StatusAccountError    = 4 //用户名或密码错误
	StatusIdentityExpired = 5 //身份过期
	StatusQueryDataError  = 6 //查询数据异常
)
