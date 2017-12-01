// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

import "testing"

func Test_GetTagsByIDS(t *testing.T) {
	tags, _ := GetTagsByIDS("1,2,33,55")
	t.Fatal(tags)
}

func Test_GetTagMsg(t *testing.T) {
	tag, _ := GetTagMsg(1)
	tag, _ = GetTagMsg(2)
	tag, _ = GetTagMsg(1)
	if tag.Type == "cate" {
		t.Fatal(tag)
	}
	t.Fatal(tag.Type)
}
