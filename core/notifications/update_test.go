// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notifications

import (
	"fmt"
	"testing"
)

func Test_Notice(t *testing.T) {
	toUser := "o7UTkjr7if4AQgcPmveQ5wJ5alsA"
	bookName := "龙符"
	chapter := "第1102章 重生"
	url := "http://readfollow.com/s/1"
	msgID, _ := Update(toUser, bookName, chapter, url)
	fmt.Println(msgID)
}
