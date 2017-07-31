// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

import (
	"fmt"
	"testing"

	"github.com/yizenghui/readfollow/conf"
	"github.com/yizenghui/readfollow/core/event"
	"github.com/yizenghui/readfollow/model"
)

func init() {
	conf.InitConfig("../conf/conf.toml")
}

func Test_GetNewBook(t *testing.T) {
	data := GetNewBook("")
	fmt.Println(data)
}

func Test_GetBookInfo(t *testing.T) {
	data, _ := GetBookInfo(1, "op")
	fmt.Println(data)
}

func Test_FindBook(t *testing.T) {
	data, err := FindBook("http://book.zongheng.com/book/652613.html")
	fmt.Println(data, err)
}

func Test_NoticeFollow(t *testing.T) {
	var book model.Book
	book.GetBookByID(1)
	users := book.GetFollowUser()
	if users != nil {
		event.BookUpdateNotice(book, users)
	}
}
