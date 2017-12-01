package orm

import "time"

// Tag 标签属性 (不存在相同的标签)
type Tag struct {
	ID        uint `gorm:"primary_key"`
	Pid       uint
	IsShow    int
	Type      string
	Title     string
	Name      string `gorm:"type:varchar(100);unique_index"`
	Intro     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// GetAllTags 获取所有 Tags
func (tag *Tag) GetAllTags() (tags []Tag) {
	DB().Find(&tags)
	return
}

// GetTagByID 获取 Tag
func (tag *Tag) GetTagByID(id int) {
	DB().First(tag, id)
}

// GetTagByName 通过名词获取标签
func (tag *Tag) GetTagByName(name string) {
	DB().Where(Tag{Name: name}).FirstOrCreate(tag)
}

// Save 保存用户信息
func (tag *Tag) Save() {
	DB().Save(&tag)
}

// GetTagsByType 通过属性获取标签
func (tag *Tag) GetTagsByType(name string) (tags []Tag) {
	DB().Where(Tag{Type: name, IsShow: 1}).Order("id ASC").Find(&tags)
	return
}

// GetTagsByTitle 通过属性获取标签
func (tag *Tag) GetTagsByTitle(name string) (tags []Tag) {
	DB().Where(Tag{Title: name}).Order("id ASC").Find(&tags)
	return
}

// GetTagsByIDS 通过ID获取标签
func (tag *Tag) GetTagsByIDS(ids []int64) (tags []Tag) {
	DB().Where(ids).Find(&tags)
	return
}
