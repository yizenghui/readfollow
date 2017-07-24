package repository

import (
	"fmt"

	"errors"

	"github.com/yizenghui/readfollow/core/common"
	"github.com/yizenghui/readfollow/model"
)

// User 所用数据包
type User struct {
	UserID    uint
	OpenID    string
	Nickname  string
	Head      string
	Books     []Book
	NotFollow bool
	SeoTag
}

//GetUser 获取最新更新书籍
func GetUser(id int, openID string) (User, error) {
	data := User{}

	var user model.User
	user.GetUserByID(id)
	if user.ID == 0 {
		return data, errors.New("我们擦身而过，不留一丝痕迹。")
	}

	data.Nickname = user.Nickname

	books := user.GetUserFollowBooks()

	if books != nil {

		for _, b := range books {
			dbo := Book{ID: b.ID, Name: b.Name, Chapter: b.Chapter, UpdatedAt: b.UpdatedAt}
			if openID != "" {
				dbo.URL = fmt.Sprintf("/s/%d?open_id=%v", b.ID, openID)
				// TODO 细分 open_id 与 uid 是否同一个人，分设书籍关注状态 (关注接口也需要做重定向)
				// if openID == user.OpenID {
				// 	dbo.UnFollowBtm = true
				// 	dbo.UnFollowLink = fmt.Sprintf("/unfollow/%d?open_id=%v", b.ID, openID)
				// }
			} else {
				dbo.URL = fmt.Sprintf("/s/%d", b.ID)
			}
			dbo.Posted = common.TransformBookPosted(b.BookURL)
			dbo.BookURL = common.TransformBookURL(b.BookURL)
			data.Books = append(data.Books, dbo)
		}
	} else {
		data.NotFollow = true
	}

	if openID != "" {
		if user.OpenID == openID {

		}
	}
	data.Title = fmt.Sprintf("%v正在跟读", data.Nickname)
	data.Description = fmt.Sprintf("%v跟读的书籍。", data.Nickname)
	data.Keywords = fmt.Sprintf("%v,跟读", data.Nickname)

	return data, nil
}

// FollowBook 关注一本书
func FollowBook(id int, openID string) (string, error) {
	var user model.User
	var book model.Book
	user.GetUserByOpenID(openID)
	if user.ID == 0 {
		return "/404.html", errors.New("我们连擦身而过的机会都不曾拥有")
	}
	book.GetBookByID(id)
	if book.ID == 0 {
		return "/404.html", errors.New("我们擦身而过，不留一丝痕迹。")
	}
	// 关注书籍
	user.UserFollowBook(book)
	return fmt.Sprintf("/s/%d?open_id=%v", id, openID), nil
}

// UnfollowBook 取消关注一本书
func UnfollowBook(id int, openID string) (string, error) {
	var user model.User
	var book model.Book
	user.GetUserByOpenID(openID)
	if user.ID == 0 {
		return "/404.html", errors.New("我们连擦身而过的机会都不曾拥有")
	}
	book.GetBookByID(id)
	if book.ID == 0 {
		return "/404.html", errors.New("我们擦身而过，不留一丝痕迹。")
	}
	// 关注书籍
	user.UserRemoveFollowBook(book)
	return fmt.Sprintf("/s/%d?open_id=%v", id, openID), nil
}
