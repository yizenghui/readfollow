// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/jinzhu/gorm"

	"fmt"
)

func Test_GetUserByOpenID(t *testing.T) {
	var db *gorm.DB
	var err error
	db, err = gorm.Open("sqlite3", "book.db")
	if err != nil {
		panic("连接数据库失败")
	}
	tx := db.Begin()
	var book Book
	// tx.Set("gorm:query_option", "FOR UPDATE").Where("publish_at > 0").First(&book)
	tx.Where("publish_at > 0").First(&book)
	fmt.Println(book)
	if book.ID > 0 {
		book.PublishAt = time.Now().Unix()

		if book.Chapter != "" && book.ChapterURL != "" {
			// client := rpc.NewClient("http://47.92.130.14:819/")

			postBook := TransformBook(book)
			if jsonStr, err := json.Marshal(postBook); err == nil {
				s, err := stub.Save(string(jsonStr))
				if err != nil {
					book.PublishAt = -1
					fmt.Println(err)
				} else {
					fmt.Println("ss:", s)
				}
			}
		}
		tx.Save(&book)
	} else {
		fmt.Println("nil")
	}
	tx.Commit()

}
