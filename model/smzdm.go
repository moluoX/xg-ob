package model

import (
	"time"
)

//SmzdmArticlePage smzdm Article page
type SmzdmArticlePage struct {
	ArticleList []SmzdmArticle `json:"article_list"`
}

//SmzdmArticle smzdm Article
type SmzdmArticle struct {
	Id         int64  `json:"article_id" xorm:"pk"`
	Title      string `json:"article_title"`
	Time       time.Time
	TimeSort   int64  `json:"article_timesort" xorm:"-"`
	Price      string `json:"article_price"`
	Unworthy   int    `json:"article_unworthy"`
	Worthy     int    `json:"article_worthy"`
	Collection int    `json:"article_collection"`
	Comment    int    `json:"article_comment"`
	Content    string `json:"article_content"`
	URL        string `json:"article_url"`
	PicURL     string `json:"article_pic_url"`
	Mall       string `json:"article_mall"`
	Link       string `json:"article_link"`
}

//SmzdmArticlePaged Smzdm Article Paged
type SmzdmArticlePaged struct {
	Rows  []SmzdmArticle `json:"rows"`
	Total int64          `json:"total"`
}
