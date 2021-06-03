package models

type Shixingzuo struct {
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

func (Shixingzuo) TableName() string {
	return "shi_xingzuo"
}

func (x *Shixingzuo) AddData() {
	DB.Create(&x)
}
