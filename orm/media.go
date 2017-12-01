package orm

import (
	"time"

	"github.com/lib/pq"
)

// Media 微信公众号
type Media struct {
	ID        uint   `gorm:"primary_key"`
	AppID     string `gorm:"type:varchar(100);unique_index"`
	AppName   string
	Cover     string
	Intro     string
	State     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time    `sql:"index"`
	Tags      pq.Int64Array `gorm:"type:int[]"` // 标签(文章继承)
	Articles  []Article
}

// GetMediaByID 获取 Media
func (media *Media) GetMediaByID(id int) {
	DB().First(media, id)
}

// GetMediaByAppID appID 获取公众号
func (media *Media) GetMediaByAppID(appID string) {
	DB().Where(Media{AppID: appID}).FirstOrCreate(media)
}

// Save 保存
func (media *Media) Save() {
	DB().Save(&media)
}
