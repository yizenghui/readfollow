// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package notifications

import (
	"fmt"
	"testing"
)

func Test_Unfollow(t *testing.T) {
	toUser := "o7UTkjr7if4AQgcPmveQ5wJ5alsA"
	bookName := "亡灵元帅"
	msgID, _ := Unfollow(toUser, bookName)
	fmt.Println(msgID)
}
