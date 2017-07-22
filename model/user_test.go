// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package model

import (
	"testing"

	"fmt"

	"github.com/yizenghui/readfollow/conf"
)

func init() {
	conf.InitConfig("../conf/conf.toml")
}

func Test_GetUser(t *testing.T) {
	var user User
	DB().First(&user, 1)
	// user.GetUserFollowBooks()
	fmt.Println(user)
}

func Test_GetUserByOpenID(t *testing.T) {
	var user User
	user.GetUserByOpenID("sss")
	// user.GetUserFollowBooks()
	DB().Save(&user)
	fmt.Println(user)

}

// Test_UserRemoveFollowBook 测试用户取消关注一本书
func Test_UserRemoveFollowBook(t *testing.T) {
	var user User
	DB().First(&user, 1)
	var book Book
	DB().First(&book, 1)
	user.UserRemoveFollowBook(book)
}

func Test_GetUserFollowBooks(t *testing.T) {
	var user User
	DB().First(&user, 1)

	books := user.GetUserFollowBooks()
	fmt.Println(books)
}

func Test_UserRemoveFollowBooks(t *testing.T) {
	var user User
	var books []Book
	DB().Limit(10).Find(&books)
	// fmt.Println(books)

	DB().First(&user, 1)

	// var book1 Book
	// var book2 Book
	// var book3 Book

	// DB().First(&book1, 1)
	// DB().First(&book2, 2)
	// DB().First(&book3, 3)

	// 测试无法删除(批量)
	// DB().Model(&user).Association("books").Delete([]Book{book1, book2, book3})

	// 测试无法删除(批量)
	// DB().Model(&user).Association("books").Delete(book1, book2, book3)

	// DB().Model(&user).Association("books").Delete(book1)

	user.UserRemoveFollowBooks(books)
}
func Test_UserRemoveAllFollow(t *testing.T) {
	var user User
	DB().First(&user, 1)
	user.UserRemoveAllFollow()
}

func Test_UserFollowBook(t *testing.T) {
	var user User
	DB().First(&user, 1)
	var book Book
	DB().First(&book, 1)
	user.UserFollowBook(book)
}

func Test_UserFollowBooks(t *testing.T) {
	var user User
	DB().First(&user, 1)
	var books []Book
	DB().Limit(10).Find(&books)
	user.UserFollowBooks(books)
}
