package common

import (
	"fmt"
	"regexp"

	"github.com/yizenghui/spider/code"
)

//TransformBookURL 把起点、横纵、17K的书籍地址转成手机版的
func TransformBookURL(url string) string {
	// 起点详细
	repQiDian := `\/\/book.qidian.com\/info\/(?P<book_id>\d+)`
	checkLinkIsQiDian, _ := regexp.MatchString(repQiDian, url)
	if checkLinkIsQiDian {
		Map := code.SelectString(repQiDian, url)
		return fmt.Sprintf("http://m.qidian.com/book/%v?from=readfollow", Map["book_id"])
	}

	// 纵横男生网
	repZongHeng := `\/\/book.zongheng.com\/book\/(?P<book_id>\d+).html`
	checkLinkIsZongHeng, _ := regexp.MatchString(repZongHeng, url)
	if checkLinkIsZongHeng {
		Map := code.SelectString(repZongHeng, url)
		return fmt.Sprintf("http://m.zongheng.com/h5/book?bookid=%v&from=readfollow", Map["book_id"])
	}

	//17K
	repSeventeenK := `www.17k.com\/book\/(?P<book_id>\w+).html`
	checkLinkIsSeventeenK, _ := regexp.MatchString(repSeventeenK, url)
	if checkLinkIsSeventeenK {
		Map := code.SelectString(repSeventeenK, url)
		return fmt.Sprintf("http://h5.17k.com/book/%v.html?from=readfollow", Map["book_id"])
	}
	return url
}

// TransformBookPosted 获取书籍首发平台
func TransformBookPosted(url string) string {
	// 起点详细
	repQiDian := `\/\/book.qidian.com\/info\/(?P<book_id>\d+)`
	checkLinkIsQiDian, _ := regexp.MatchString(repQiDian, url)
	if checkLinkIsQiDian {
		return "qidian.com"
	}

	// 纵横男生网
	repZongHeng := `\/\/book.zongheng.com\/book\/(?P<book_id>\d+).html`
	checkLinkIsZongHeng, _ := regexp.MatchString(repZongHeng, url)
	if checkLinkIsZongHeng {
		return "zongheng.com"
	}

	//17K
	repSeventeenK := `www.17k.com\/book\/(?P<book_id>\w+).html`
	checkLinkIsSeventeenK, _ := regexp.MatchString(repSeventeenK, url)
	if checkLinkIsSeventeenK {
		return "17k.com"
	}
	return "未知平台"
}
