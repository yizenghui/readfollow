package repository

import (
	"fmt"
	"strconv"

	"github.com/yizenghui/readfollow/core/common"
	"github.com/yizenghui/readfollow/core/event"
	"github.com/yizenghui/sda/code"
)

// SignTask 签名任务数据包
type SignTask struct {
	ID        int32
	Ticket    string
	QrcodeURL string
	Callback  string
}

// CreateWebGetSignTask 站点获取签名任务
func CreateWebGetSignTask(callback string) (SignTask, error) {
	var st SignTask
	st.ID = common.CreateTask()
	qrcode, err := event.CreateTempQrcode(st.ID)
	if err != nil {
		return st, err
	}
	st.Ticket = qrcode.Ticket
	st.QrcodeURL = fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%v", st.Ticket)
	st.Callback = callback
	common.SetTicket(st.ID, st.Ticket)
	return st, nil
	// return st
}

// SignTaskValue 签名任务数据包
type SignTaskValue struct {
	ID    int32
	Value string
}

//GetWebGetSignTaskValue 获取签名值，没有返回空
func GetWebGetSignTaskValue(id int32) (SignTaskValue, error) {
	stv := SignTaskValue{}
	val, err := common.GetTaskCompleteValue(id)
	if err != nil {
		return stv, err
	}
	stv.ID = id
	stv.Value = val
	return stv, nil
}

//SetWebGetSignTaskValue 设置签名值
func SetWebGetSignTaskValue(id int32, openID string) bool {

	SetTaskValue(id, openID)
	return true

}

//RemoveSignTask 移除任务
func RemoveSignTask(id int32) {
	common.DeleteTask(id)
}

// SetWebGetSignTaskValueForWechatPush ..
func SetWebGetSignTaskValueForWechatPush(str, openID string) bool {
	i64, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		// 这里面用正则匹配出整数
		istr := code.FindString(`(?P<int>\d+)`, str, "int")
		i64, _ = strconv.ParseInt(istr, 10, 64)
	}
	return SetWebGetSignTaskValue(int32(i64), openID)
}

//SetTaskValue 设置签名值
func SetTaskValue(id int32, value string) {
	common.SetValue(id, value)
}
