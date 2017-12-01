// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

import "testing"

func Test_GetSign(t *testing.T) {
	js, err := GetSign("http://readfoww.com/t/1")
	if err != nil {
		panic(err)
	}
	t.Fatal(js)
}
