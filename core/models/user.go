package models

import "github.com/jinzhu/gorm"

// User has and belongs to many languages, use `user_languages` as join table
type User struct {
	gorm.Model
	OpenID   string `gorm:"size:255"`
	Nickname string `gorm:"size:255"`
	Head     string `gorm:"size:255"`
	Books    []Book `gorm:"many2many:user_books;"` // 用户关注的书
}

func GetUserForBook() {
	//o7UTkjr7if4AQgcPmveQ5wJ5alsA

}
