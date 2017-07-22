// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

import (
	"testing"

	"fmt"

	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/menu"
	"github.com/chanxuehong/wechat.v2/mp/user"
)

func Test_GetMenu(t *testing.T) {
	var wxAppID = "wx702b93aef72f3549"
	var wxAppSecret = "8b69f45fc737a938cbaaffc05b192394"
	ats := core.NewDefaultAccessTokenServer(wxAppID, wxAppSecret, nil)
	clt := core.NewClient(ats, nil)
	m, p, _ := menu.Get(clt)
	fmt.Println(m, p)

	var buttons []menu.Button
	but1 := menu.Button{
		Type: "click",
		Name: "我的关注",
		Key:  "myfollow",
	}

	buttons = append(buttons, but1)
	fmt.Println(buttons)
	// newMenu := &menu.Menu{
	// 	Buttons: buttons,
	// }
	// menu.Create(clt, newMenu)
	// menu.Menu{}
	// OpenID := "xsasxa"
	// var wxAppID = "wx702b93aef72f3549"
	// var wxAppSecret = "8b69f45fc737a938cbaaffc05b192394"
	// ats := core.NewDefaultAccessTokenServer(wxAppID, wxAppSecret, nil)
	// clt := core.NewClient(ats, nil)
	// // user,err := user.Get(clt,OpenID,"zh_CN")
	// user, _ := user.Get(clt, OpenID, "zh_CN")
	// fmt.Println(user)
}

func Test_GetUserInfo(t *testing.T) {
	OpenID := "o7UTkjr7if4AQgcPmveQ5wJ5alsA"
	var wxAppID = "wx702b93aef72f3549"
	var wxAppSecret = "8b69f45fc737a938cbaaffc05b192394"
	ats := core.NewDefaultAccessTokenServer(wxAppID, wxAppSecret, nil)
	clt := core.NewClient(ats, nil)
	// user,err := user.Get(clt,OpenID,"zh_CN")
	user, _ := user.Get(clt, OpenID, "zh_CN")
	fmt.Println(user)
}
