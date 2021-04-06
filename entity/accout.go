package entity

type Login struct{
	Id int64 `json:"id"`
	User_Id string `json:"userId"`
	Account string `json:"account"`
	Password string `json:"password"`
	UserToken string `json:"userToken"`
	DeviceToken string `json:"deviceToken"`
}