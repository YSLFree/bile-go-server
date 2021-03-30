package users_info

// ResultData 返回jie guo
type ResultData struct {
	Code  int32 //返回错误码
	State bool  //false 失败，ture 成功
	Data  interface{}
}

//LoginInfo 账号密码
type LoginInfo struct {
	Account  string // 账号
	Password string ///密码
}

//UserInfos 用户详细信息
type UserInfos struct {
	Userid            int32  //用户id
	Name              string //姓名
	Nickname          string //昵称
	Gender            int8   //性别 0女，1男
	Age               int8   //年龄
	Signature         string //签名
	Head              string //头像地址
	Constellation     int8   //星座
	Locaion           string //居住地
	Hometown          string //家乡
	Height            int16  //身高
	Weight            int16  //体重
	Profession        string //职业
	EmotionalState    int8   //情感状态 ,0单身
	Isforbindden      int8   //用户是否被封杀 0 没有被封杀 1被封杀
	Level             int8   //用户等级
	Fans              int32  //粉丝
	Fcous             int32  //关注
	Friends           int32  //好友数
	Labels            string //标签
	Covers            string //封面
	Favorite          string //爱好
	IntroductionVideo string //介绍视频
	IntroductionAudio string //介绍语音
	Receivedgift      Gifts  //收到的礼物
}

//vip
type Vip struct {
	IsVip     bool //是否是vip
	VipLeavel int8 //vip等级
}

//收到的礼物
type Gifts struct {
	Giftid   int8   //礼物id
	Giftname string //礼物名字
	Gifturl  string //礼物url
}

//Dynamic 动态
type Dynamic struct {
	Dynamicid    int64  //动态id
	Userid       int32  //用户id
	Username     string //用户名
	Gender       int8   //性别：0女，1男
	Headurl      string //头像
	Age          int8   //年龄
	Isminefcous  int8   //是否是自己关注的动态 0未关注 ，1关注
	Showtype     int8   // 帆布类型： 0纯文字，1纯图片，2纯视频，3文字+图片，4文字+视频
	Showconent   string // 文字内容
	Showvideo    string //视频地址
	Showimage    string // 图片地址
	Showtime     int64  //发布时间
	Viewcount    int32  //浏览次数
	Showlocation string //发布位置
	Addcount     int32  //点赞数
	Discusscount int32  //评论数

}

//DynamicDiscuss 朋友圈评论
type DynamicDiscuss struct {
	Dynamicid      int64  //动态id
	Userid         int32  // 用户id
	Username       string //用户名
	Discusscontent string //评论内容
}
