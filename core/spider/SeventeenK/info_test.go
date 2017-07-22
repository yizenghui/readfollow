// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package SeventeenK

import (
	"fmt"
	"testing"
)

func Test_GetInfo(t *testing.T) {
	// http://www.17k.com/book/2469390.html
	//http://www.17k.com/book/777148.html
	book, err := GetInfo("http://www.17k.com/book/777148.html")

	if err != nil {
		panic("spider data error")
	}
	fmt.Println(book)

	// fmt.Printf(" %v %v-> %v %v %v  %v %v %v \n", book.IsVIP, book.Date, book.Name, book.Author, book.Chapter, book.ChapterURL, book.AuthorURL, book.BookURL)

}
