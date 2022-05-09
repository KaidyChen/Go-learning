package models

type Student struct {
	Id int `json:"id"`
	Number string `json:"number"`
	Password string `json:"password"`
	ClassId int `json:"class_id"`
	Name string `json:"name"`
	Lesson []Lesson `gorm:"many2many:lesson_student;" json:"lesson"`

}

func (Student) TableName() string {
	return "student"
}
