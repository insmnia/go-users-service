package service

type User struct {
	ID       uint   `gorm:"column:id;primary_key"`
	Username string `gorm:"column:username;unique_index"`
	Password string `gorm:"column:password;not_null"`
}
