package models

import "time"

//内存对齐
type Post struct {
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	ID          int64     `json:"id" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id" binding:"required"`
	CommunityID int64     `json:"community_id"  db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

//帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName       string `json:"author_name"`
	*Post                   //嵌入帖子
	*CommunityDetail        //嵌入社区信息
}
