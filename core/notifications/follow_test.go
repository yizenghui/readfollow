// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notifications

import (
	"fmt"
	"testing"
)

func Test_Follow(t *testing.T) {
	toUser := "o7UTkjr7if4AQgcPmveQ5wJ5alsA"
	bookName := "亡灵元帅"
	// 通知用户已经关注XX小说，当该小说更新时会提醒用户查看
	msgID, _ := Follow(toUser, bookName)
	fmt.Println(msgID)
}
