package event

import (
	"fmt"

	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/qrcode"
)

const wxAppID = "wx702b93aef72f3549"
const wxAppSecret = "8b69f45fc737a938cbaaffc05b192394"

//CreateTempQrcode 创建临时二维码
func CreateTempQrcode(id int32) {

	ats := core.NewDefaultAccessTokenServer(wxAppID, wxAppSecret, nil)
	clt := core.NewClient(ats, nil)

	q, e := qrcode.CreateTempQrcode(clt, id, 7200)

	if e != nil {

	}
	fmt.Println(q)

}
