// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

import (
	"testing"

	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Test_TransformURL(t *testing.T) {
	// 17K
	u1 := TransformBookURL("http://www.17k.com/book/2124315.html")
	p1 := TransformBookPosted("http://www.17k.com/book/2124315.html")
	// 纵横
	u2 := TransformBookURL("http://book.zongheng.com/book/683149.html")
	p2 := TransformBookPosted("http://book.zongheng.com/book/683149.html")
	// 起点
	u3 := TransformBookURL("http://book.qidian.com/info/1009961125")
	p3 := TransformBookPosted("http://book.qidian.com/info/1009961125")

	fmt.Println(u1, p1, u2, p2, u3, p3)

}
