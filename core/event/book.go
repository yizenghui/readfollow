package event

import (
	"fmt"

	"github.com/yizenghui/readfollow/conf"
	"github.com/yizenghui/readfollow/core/notifications"
	"github.com/yizenghui/readfollow/model"
)

//BookUpdateNotice 更新提醒
func BookUpdateNotice(book model.Book, users []model.User) {

	for _, user := range users {
		url := fmt.Sprintf("http://%v/s/%d?open_id=%v", conf.Conf.App.Host, book.ID, user.OpenID)
		notifications.Update(user.OpenID, book.Name, book.Chapter, url)
	}

}
