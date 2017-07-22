// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package conf

import (
	"fmt"
	"testing"
)

func init() {

	err := InitConfig("conf.toml")
	fmt.Println(err)
}
func Test_GetFans(t *testing.T) {
	fmt.Println(Conf.App.Name)
	fmt.Println(Conf.Wechat.OriID, Conf.Wechat.AppID, Conf.Wechat.Token, Conf.Wechat.AesKey)
}
