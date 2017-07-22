// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"testing"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Test_InitUser(t *testing.T) {

	var db *gorm.DB

	var err error
	// db, err = gorm.Open("sqlite3", "book.db")
	db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=spider sslmode=disable password=123456")

	if err != nil {
		panic("连接数据库失败")
	}

	// 自动迁移模式
	db.AutoMigrate(&User{})
	defer db.Close()

	var user1 User
	db.Where(User{OpenID: "o7UTkjr7if4AQgcPmveQ5wJ5alsA"}).FirstOrCreate(&user1)

	if user1.Nickname != "yize" {
		user1.Nickname = "yize"
		db.Save(&user1)
	}

	var user2 User
	db.Where(User{OpenID: "damingdamingdamingdamingdaming"}).FirstOrCreate(&user2)

	if user2.Nickname != "daming" {
		user2.Nickname = "daming"
		db.Save(&user2)
	}

	var user3 User
	db.Where(User{OpenID: "testtesttesttesttesttesttest"}).FirstOrCreate(&user3)

	if user3.Nickname != "test" {
		user3.Nickname = "test"
		db.Save(&user3)
	}

}

func Test_InitUserFollow(t *testing.T) {

	var db *gorm.DB

	var err error
	// db, err = gorm.Open("sqlite3", "book.db")
	db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=spider sslmode=disable password=123456")

	if err != nil {
		panic("连接数据库失败")
	}

	// 自动迁移模式
	db.AutoMigrate(&User{}, &Book{})
	defer db.Close()

	var user1 User
	db.First(&user1, 1)

	var user2 User
	db.First(&user2, 2)

	//TODO 需要验证地址是否会改变
	// 章节地址与数据库中的不同
	// if user.Nickname != "xiaoyi" {
	// 	user.Nickname = "xiaoyi"
	// 	db.Save(&user)
	// }

	var book1 Book
	db.First(&book1, 12)
	db.Model(&user1).Association("books").Append(book1)

	var book2 Book
	db.First(&book2, 22)
	db.Model(&user1).Association("books").Append(book2)
	db.Model(&user2).Association("books").Append(book2)

	var book3 Book
	db.First(&book3, 333)
	db.Model(&user1).Association("books").Append(book3)

}

func Test_UserFollowBook(t *testing.T) {

	var db *gorm.DB

	var err error
	// db, err = gorm.Open("sqlite3", "book.db")
	db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=spider sslmode=disable password=123456")

	if err != nil {
		panic("连接数据库失败")
	}

	// 自动迁移模式
	db.AutoMigrate(&User{})
	defer db.Close()

	var user User
	db.First(&user, 1)

	var book Book
	db.First(&book, 12)
	db.Model(&user).Association("books").Append(book)

}

func Test_UserFollowCount(t *testing.T) {

	var db *gorm.DB

	var err error
	// db, err = gorm.Open("sqlite3", "book.db")
	db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=spider sslmode=disable password=123456")

	if err != nil {
		panic("连接数据库失败")
	}

	// 自动迁移模式
	db.AutoMigrate(&User{})
	defer db.Close()

	var user User
	db.First(&user, 1)

	// var books []Book
	// db.Find(&books)
	// db.Model(&user).Association("books").Append(books)

	count := db.Model(&user).Where("book_id = ?", 12).Association("books").Count()
	fmt.Println(count)
}

func Test_UserFollowAllBook(t *testing.T) {

	var db *gorm.DB

	var err error
	// db, err = gorm.Open("sqlite3", "book.db")
	db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=spider sslmode=disable password=123456")

	if err != nil {
		panic("连接数据库失败")
	}

	// 自动迁移模式
	db.AutoMigrate(&User{})
	defer db.Close()

	var user User
	db.First(&user, 1)

	var books []Book
	db.Find(&books)
	db.Model(&user).Association("books").Append(books)

}

func Test_UserUnFollowBook(t *testing.T) {

	var db *gorm.DB

	var err error
	// db, err = gorm.Open("sqlite3", "book.db")
	db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=spider sslmode=disable password=123456")

	if err != nil {
		panic("连接数据库失败")
	}

	// 自动迁移模式
	db.AutoMigrate(&User{})
	defer db.Close()

	var user User
	db.First(&user, 1)

	var book Book
	db.First(&book, 12)
	// 移除指定关联关系
	// db.Model(&user).Association("books").Delete(book)

	// 移除所有关联关系
	db.Model(&user).Association("books").Clear()

}

func Test_UserFollowBooks(t *testing.T) {

	var db *gorm.DB

	var err error
	// db, err = gorm.Open("sqlite3", "book.db")
	db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=spider sslmode=disable password=123456")

	if err != nil {
		panic("连接数据库失败")
	}

	// 自动迁移模式
	db.AutoMigrate(&User{})
	defer db.Close()

	var user User
	db.First(&user, 1)

	var books []Book

	db.Model(&user).Association("books").Find(&books)

	fmt.Println(books)
	// 移除所有关联关系
	// db.Model(&user).Association("books").Clear()

}

func Test_BookFollows(t *testing.T) {

	var db *gorm.DB

	var err error
	// db, err = gorm.Open("sqlite3", "book.db")
	db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=spider sslmode=disable password=123456")

	if err != nil {
		panic("连接数据库失败")
	}

	// 自动迁移模式
	db.AutoMigrate(&User{}, &Book{})
	defer db.Close()

	var book Book
	db.First(&book, 22)

	var users []User

	db.Model(&book).Association("users").Find(&users)

	fmt.Println(users)
	// 移除所有关联关系
	// db.Model(&user).Association("books").Clear()

}
