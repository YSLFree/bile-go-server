package nets

const (
	//OkResultCode 成功
	OkResultCode = 0x000000
	//ParamsFormErrorFormClient 服务器接收客户端参数解析错误
	ParamsFormErrorFormClient = 0x000001
	//AccountOrPasswordError 账号或密码错误
	AccountOrPasswordError = 0x0000002
	//MysqlConnectError 数据库连接失败
	MysqlConnectError = 0x0000003
	//SqlActionError 数据库操作失败
	SqlActionError = 0x00000004
	//UserExist 用户已经存在
	UserExist = 0x0000005
	//AccountFormatError 账号格式不正确
	AccountFormatError = 0x0000006
)
