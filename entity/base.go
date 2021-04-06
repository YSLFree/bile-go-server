package entity

//基本数据
type BaseEntity struct {
	Code  int16       `json:"code"`
	State bool        `json:"state"`
	Data  interface{} `json:"data"`
}
