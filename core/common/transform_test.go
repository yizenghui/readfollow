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

func Test_TransformData(t *testing.T) {
	arr := []string{
		"12万",
		"12.34万",
		"12.34",
		"12.00",
		"123446",
		"1234",
	}

	for _, k := range arr {
		i := TransformBookTotal(k)
		f := PrintBookTotal(i)
		fmt.Println(i, f)
	}

}
