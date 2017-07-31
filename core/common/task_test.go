// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

import (
	"testing"

	"fmt"
)

func init() {
	// t1 := CreateTask()
	// t2 := CreateTask()
	// t3 := CreateTask()
	// fmt.Println(t1, t2, t3)
	// SetValue(100002, "bbb")
}

func Test_CreateSignTask(t *testing.T) {

	SetValue(100001, "a")
	SetValue(100001, "aaa")

	DeleteTask(3)

	t1, _ := GetTask(100001)
	t2, _ := GetTask(100002)
	t3, _ := GetTask(100003)

	fmt.Println(t1, t2, t3)
}
