// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package model

import (
	"fmt"
	"testing"

	"github.com/yizenghui/readfollow/conf"
)

func Test_GetDB(t *testing.T) {
	conf.InitConfig("../conf/conf.toml")
	fmt.Println(conf.Conf.App.Name)
	fmt.Println(DB())
}
