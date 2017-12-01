package orm

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Article 文章列表
type Article struct {
	ID      uint      `gorm:"primary_key"`
	Title   string    // 标题
	Author  string    // 作者
	Cover   string    // 封面
	Intro   string    // 介绍
	PubAt   time.Time `sql:"index"` // 微信文章发布时间
	MediaID uint
	// Media     Media
	Like      int64         `gorm:"default:0"`
	Hate      int64         `gorm:"default:0"`
	View      int64         `gorm:"default:0"`                      // 点击次数，通过它进行计算排名
	URL       string        `gorm:"type:varchar(255);unique_index"` // 微信文章地址
	Rank      float64       `sql:"index"`                           // 排行
	Tags      pq.Int64Array `gorm:"type:int[]"`                     // 标签
	State     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// GetArticleByID 获取Article
func (article *Article) GetArticleByID(id int) {
	DB().Preload("Media").First(article, id)
}

// GetArticleByURL 通过url获取Article 如果没有的话进行初始化 (注：此url由文章详细页获得)
func (article *Article) GetArticleByURL(url string) {
	DB().Where(Article{URL: url}).FirstOrCreate(article)
}

// Save 保存用户信息
func (article *Article) Save() {
	DB().Save(&article)
}

// Hot 热门
func (article *Article) Hot(limit, offset int) (articles []Article) {
	DB().Offset(offset).Limit(limit).Order("rank DESC").Find(&articles)
	return
}

// New 最新
func (article *Article) New(limit, offset int) (articles []Article) {
	DB().Offset(offset).Limit(limit).Order("id DESC").Find(&articles)
	return
}

// GetArticle 获取文章
func (article *Article) GetArticle(limit, offset, tag int, order string) (articles []Article) {
	// var selectTag pq.Int64Array
	if tag != 0 {
		// selectTag = append(selectTag, int64(tag))
		//Article{Tags: selectTag}   "tags && {?}", selectTag
		DB().Preload("Media").Where("tags @> ?", fmt.Sprintf("{%d}", tag)).Offset(offset).Limit(limit).Order(order).Find(&articles)
	} else {
		DB().Preload("Media").Offset(offset).Limit(limit).Order(order).Find(&articles)
	}
	return
}

// GetArticleCursorByID 以ID为游标获取数据
func (article *Article) GetArticleCursorByID(id, limit int, tags []int64) (articles []Article) {

	DB().Scopes(ScopesCursorID(int64(id))).Preload("Media").Scopes(ScopesTag(tags)).Limit(limit).Order("id DESC").Find(&articles)

	return
}

// GetArticleCursorByRank 以Rank为游标获取数据
func (article *Article) GetArticleCursorByRank(rank float64, limit int, tags []int64) (articles []Article) {

	DB().Scopes(ScopesCursorRank(rank)).Preload("Media").Scopes(ScopesTag(tags)).Limit(limit).Order("rank DESC").Find(&articles)

	return
}

//ScopesTag 生成tag查询条件
func ScopesTag(tags []int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(tags) > 0 {
			var p []string
			for _, t := range tags {
				p = append(p, strconv.Itoa(int(t)))
			}
			str := strings.Join(p, ",")
			return db.Where("tags && ?", fmt.Sprintf("{%v}", str))
		}
		return db
	}
}

//ScopesCursorID 生成ID查询条件
func ScopesCursorID(id int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if id > 0 {
			return db.Where("id < ?", id)
		}
		return db
	}
}

//ScopesCursorRank 生成rank查询条件
func ScopesCursorRank(rank float64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if rank > 0 {
			return db.Where("rank < ?", rank)
		}
		return db
	}
}

// func Cursor(status []string) func (db *gorm.DB) *gorm.DB {
//     return func (db *gorm.DB) *gorm.DB {
//         return db.Scopes(AmountGreaterThan1000).Where("status in (?)", status)
//     }
// }

// GetHotArticleByTag 获取文章
func (article *Article) GetHotArticleByTag(limit, offset, tag int) (articles []Article) {
	var selectTag pq.Int64Array
	selectTag = append(selectTag, int64(tag))
	DB().Where(Article{Tags: selectTag}).Offset(offset).Limit(limit).Order("rank DESC").Find(&articles)
	return
}

// GetNewArticleByTag 获取文章
func (article *Article) GetNewArticleByTag(limit, offset, tag int) (articles []Article) {
	var selectTag pq.Int64Array
	selectTag = append(selectTag, int64(tag))
	DB().Where(Article{Tags: selectTag}).Offset(offset).Limit(limit).Order("id DESC").Find(&articles)
	return
}
