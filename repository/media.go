package repository

import (
	"errors"

	"github.com/GanEasy/wxrankapi/orm"
)

func GetMediaIDByAppID() {

}

//GetMediaByAppID 通过 appID 获取公众号信息
func GetMediaByAppID(appID string) (media orm.Media, err error) {

	if appID != "" {

		media.GetMediaByAppID(appID)
		return media, nil
	}
	return media, errors.New("出错了")

}

//GetMediaByID 通过 appID 获取公众号信息
func GetMediaByID(id int) (media orm.Media, err error) {

	if id > 0 {
		media.GetMediaByID(id)
		return media, nil
	}

	return media, errors.New("获取失败")

}
