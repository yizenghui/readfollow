// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package orm

import (
	"testing"

	"fmt"
)

func init() {

	// DB().AutoMigrate(&Post{})
}

func Test_GetPost(t *testing.T) {
	var Post Post
	DB().First(&Post, 1)
	// Post.GetPostFollowBooks()
	fmt.Println(Post)
}
