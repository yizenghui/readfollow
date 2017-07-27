// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

import (
	"fmt"
	"testing"

	"github.com/yizenghui/readfollow/conf"
)

func init() {
	conf.InitConfig("../conf/conf.toml")
}

func Test_CreateWebGetSignTask(t *testing.T) {
	data, err := CreateWebGetSignTask()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
