// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package event

import (
	"testing"
	"time"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/yizenghui/read-follow/core/common"
	"github.com/yizenghui/read-follow/core/models"
	"github.com/yizenghui/sda"
)

func Test_BookFollowsNotice(t *testing.T) {

	var db *gorm.DB

	var err error
	// db, err = gorm.Open("sqlite3", "book.db")
	db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=spider sslmode=disable password=123456")

	if err != nil {
		panic("连接数据库失败")
	}

	// 自动迁移模式
	db.AutoMigrate(&models.User{}, &models.Book{})
	defer db.Close()

	var book models.Book
	db.First(&book, 1)

	var users []models.User

	db.Model(&book).Association("users").Find(&users)

	BookUpdateNotice(book, users)
	// 移除所有关联关系
	// db.Model(&user).Association("books").Clear()

}

func Test_UserFollowBookForURL(t *testing.T) {
	query := "http://book.qidian.com/info/1005986994"
	openID := "o7UTkjr7if4AQgcPmveQ5wJ5alsA"
	spiderBook, _ := sda.FindBookBaseByBookURL(query)
	if spiderBook.Name != "" {

		db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=spider sslmode=disable password=123456")
		if err != nil {
			panic("连接数据库失败")
		}

		var book models.Book

		db.Where(models.Book{BookURL: spiderBook.BookURL}).FirstOrCreate(&book)

		book.Name = spiderBook.Name
		book.Author = spiderBook.Author
		book.Chapter = spiderBook.Chapter
		book.Total = spiderBook.Total
		book.AuthorURL = spiderBook.AuthorURL
		book.ChapterURL = spiderBook.ChapterURL
		book.BookURL = spiderBook.BookURL

		// TODO 获取票数
		vote := 1   // 支持
		devote := 0 // 反对
		level := 0  //级别
		// 获取排行分数
		book.Rank = common.GetRank(vote, devote, time.Now().Unix(), level)
		db.Save(&book)

		fmt.Println(book)
		if openID != "" {
			var user models.User
			db.Where("open_id = ?", openID).First(&user)
			if user.ID == 0 {
				// return c.Render(http.StatusOK, "404", "")
			} else {
				// 关注
				db.Model(&user).Association("books").Append(book)

			}
		}
	}

}
