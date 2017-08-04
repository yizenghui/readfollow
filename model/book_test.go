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

func Test_GetBook(t *testing.T) {
	var book Book
	// DB().AutoMigrate(&Book{})

	DB().First(&book, 1)
	users := book.GetFollowUser()
	fmt.Println(users)
}

func Test_GetHotBooks(t *testing.T) {
	var book Book
	books := book.GetHotBooks()
	fmt.Println(books)
}
