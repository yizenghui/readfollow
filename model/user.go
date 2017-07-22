package model

import "time"

// User has and belongs to many languages, use `user_languages` as join table
type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	OpenID    string     `gorm:"size:255"`
	Nickname  string     `gorm:"size:255"`
	Head      string     `gorm:"size:255"`
	Books     []Book     `gorm:"many2many:user_books;"` // 用户关注的书
}

// GetUserByID 获取用户关注的书籍
func (user *User) GetUserByID(id int) {
	DB().First(user, id)
}

// GetUserFollowBooks 获取用户关注的书籍
func (user *User) GetUserFollowBooks() []Book {
	var books []Book
	//o7UTkjr7if4AQgcPmveQ5wJ5alsA
	DB().Model(&user).Association("books").Find(&books)
	return books
}

// UserRemoveAllFollow 用户取消所有关注
func (user *User) UserRemoveAllFollow() {
	// 移除所有关联关系
	DB().Model(&user).Association("books").Clear()
}

// UserRemoveFollowBook 用户取消一本书
func (user *User) UserRemoveFollowBook(book Book) {
	// 移除所有关联关系
	DB().Model(&user).Association("books").Delete(book)
}

// UserRemoveFollowBooks 用户取消关注多本书
func (user *User) UserRemoveFollowBooks(books []Book) {
	// 移除所有关联关系
	// TODO 这个不行滴，不清楚怎么回事
	DB().Model(&user).Association("books").Delete(books)
}

// UserFollowBook 用户关注一本书
func (user *User) UserFollowBook(book Book) {
	// 添加一个关联关系
	DB().Model(&user).Association("books").Append(book)
}

// CheckUserIsFollowBook  检查用户是否关注某书
func (user *User) CheckUserIsFollowBook(book Book) bool {
	// 添加一个关联关系
	HasFollow := DB().Model(&user).Where("book_id = ?", book.ID).Association("books").Count()
	if HasFollow == 0 {
		return false
	}
	return true
}

// UserFollowBooks 用户关注多本书
func (user *User) UserFollowBooks(books []Book) {
	// 添加一个关联关系
	DB().Model(&user).Association("books").Append(books)
}

// GetUserByOpenID 通过openID获取用户信息 如果没有的话进行初始化
func (user *User) GetUserByOpenID(openID string) {
	DB().Where(User{OpenID: openID}).FirstOrInit(user)
}

// Save 保存用户信息
func (user *User) Save() {
	DB().Save(&user)
}
