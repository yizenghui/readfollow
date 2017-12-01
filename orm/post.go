package orm

import "time"

// Post has and belongs to many languages, use `Post_languages` as join table
type Post struct {
	ID         uint   `gorm:"primary_key"`
	URL        string `gorm:"type:varchar(255);unique_index"`
	ArticleURL string
	State      int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}

// GetPostByID 获取post
func (post *Post) GetPostByID(id int) {
	DB().First(post, id)
}

// GetPostByURL 通过url获取post 如果没有的话进行初始化 (注：此url由文章详细页获得)
func (post *Post) GetPostByURL(url string) {
	DB().Where(Post{URL: url}).FirstOrCreate(post)
}

// Save 保存用户信息
func (post *Post) Save() {
	DB().Save(&post)
}
