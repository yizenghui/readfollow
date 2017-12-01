// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

import "testing"

func Test_GetRank(t *testing.T) {
	r1 := Rank(2, 0, 1510416001)
	r2 := Rank(1, 0, 1510416001)
	t.Fatal(r1, r2)
}
