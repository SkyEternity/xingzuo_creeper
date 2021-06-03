package models

type Dyxingzuo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Author      string `json:"author"`
	Query       string `json:"query"`
	Desc        string `json:"desc"`
	Content     string `json:"content"`
	ContentTime string `json:"content_time"`
	CreateTime  MyTime `json:"create_time" gorm:"autoCreateTime"`
}

func (Dyxingzuo) TableName() string {
	return "dy_xingzuo"
}

func (x *Dyxingzuo) AddData() {
	DB.Create(&x)
}
