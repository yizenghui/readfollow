package repository

import (
	"errors"
	"strconv"
	"strings"

	"github.com/GanEasy/wxrankapi/orm"
)

// TagItem 标签属性 (不存在相同的标签)
type TagItem struct {
	ID    uint
	Type  string
	Title string
}

//CacheTags 缓存标签
var CacheTags = map[int]TagItem{}

func init() {
	var tag orm.Tag
	tags := tag.GetAllTags()

	for _, t := range tags {
		CacheTags[int(t.ID)] = TagItem{
			t.ID,
			t.Type,
			t.Title,
		}
	}

	// var initTag = []string{
	// 	"文博",
	// 	"汽车",
	// 	"IT",
	// 	"生活",
	// 	"自媒体",
	// 	"其它",
	// }

	// for _, v := range initTag {
	// 	var tag orm.Tag
	// 	tag.GetTagByName(v)
	// 	tag.Type = "cate"
	// 	tag.Title = v
	// 	tag.IsShow = 1
	// 	tag.Save()
	// }

}

//GetTagMsg ..
func GetTagMsg(id int) (tag TagItem, err error) {

	if t, ok := CacheTags[id]; ok {
		tag = t
	} else {
		newTag, err := Tag(id)
		if err == nil {
			tag.ID = newTag.ID
			tag.Title = newTag.Title
			tag.Type = newTag.Type
			CacheTags[int(newTag.ID)] = tag
		}
	}

	return
}

//Tag ..
func Tag(id int) (tag orm.Tag, err error) {

	// var a orm.Article
	tag.GetTagByID(id)

	if tag.Title == "" {
		err = errors.New("内容异常")
		return
	}

	return
}

//GetTagByMediaID ..
func GetTagByMediaID(id int) (tag orm.Tag, err error) {
	media, err := GetMediaByID(id)
	if media.AppID != "" {
		// var a orm.Article
		tag.GetTagByName(media.AppID)
	} else {

		err = errors.New("无法获取公众号AppID")
	}

	if tag.Title == "" {
		err = errors.New("获取标签失败")
	}

	return
}

//GetTagByType ..通过属性获取标签
func GetTagByType(name string) (tags []orm.Tag, err error) {
	if name != "" {
		var tag orm.Tag
		tags = tag.GetTagsByType(name)
	}
	return
}

//GetTagsByTitle ..通过属性获取标签
func GetTagsByTitle(name string) (tags []orm.Tag, err error) {
	if name != "" {
		var tag orm.Tag
		tags = tag.GetTagsByTitle(name)
	}
	return
}

//GetTagsByIDS ..通过id获取标签
func GetTagsByIDS(idstr string) (tags []orm.Tag, err error) {
	var ids []int64
	arr := strings.Split(idstr, ",")
	for _, id := range arr {
		c, e := strconv.Atoi(id)
		if e == nil {
			ids = append(ids, int64(c))
		}
	}
	if len(ids) > 0 {
		var tag orm.Tag
		tags = tag.GetTagsByIDS(ids)
	}
	return
}
