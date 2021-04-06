package entity

//UserInfos 用户详细信息
type UserInfos struct {
	Id                int    `json:"id"`                //表id
	User_Id           string `json:"usdrId"`            //用户id
	Name              string `json:"name"`              //姓名
	NickName          string `json:"nickname"`          //昵称
	Gender            int8   `json:"gender"`            //性别 0女，1男
	Age               int8   `json:"age"`               //年龄
	Signature         string `json:"signature"`         //签名
	HeadUrl           string `json:"head"`              //头像地址
	Constellation     int8   `json:"constellation"`     //星座
	Location          string `json:"location"`          //居住地
	Hometown          string `json:"hometown"`          //家乡
	Height            int16  `json:"height"`            //身高
	Weight            int16  `json:"weight"`            //体重
	Profession        string `json:"profession"`        //职业
	EmotionalState    int8   `json:"emotionalState"`    //情感状态 ,0单身
	Isforbindden      int8   `json:"isForbindden"`      //用户是否被封杀 0 没有被封杀 1被封杀
	Level             int8   `json:"level"`             //用户等级
	Fans              int32  `json:"fans"`              //粉丝
	Fcous             int32  `json:"fcous"`             //关注
	Friends           int32  `json:"friends"`           //好友数
	Labels            string `json:"labels"`            //标签
	Covers            string `json:"covers"`            //封面
	Favorite          string `json:"favorite"`          //爱好
	IntroductionVideo string `json:"introductionVideo"` //介绍视频
	IntroductionAudio string `json:"introductionAudio"` //介绍语音
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
