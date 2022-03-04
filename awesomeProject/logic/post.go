package logic

import (
	"fmt"
	"golangstudy/jike/awesomeProject/dao/mysql"
	"golangstudy/jike/awesomeProject/dao/redis"
	"golangstudy/jike/awesomeProject/models"
	"golangstudy/jike/awesomeProject/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	p.ID = snowflake.GenID()

	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID)
	return
}

func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {

	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql get post id failed", zap.Error(err))
		return
	}
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql get user id failed", zap.Error(err))
		return
	}
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		fmt.Println(post.CommunityID)
		zap.L().Error("mysql get community id failed", zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}

	return
}
func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}

	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql get user id failed", zap.Error(err))
			continue
		}
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			fmt.Println(post.CommunityID)
			zap.L().Error("mysql get community id failed", zap.Error(err))
			continue
		}

		postdetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postdetail)
	}

	return
}
