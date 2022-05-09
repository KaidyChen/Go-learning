package models

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Age int `json:"age"`
	Email string `json:"email"`
	AddTime int `json:"add_time"`
}

//定义结构体操作的数据库表
func (User) TableName() string {
	return "user"
}
