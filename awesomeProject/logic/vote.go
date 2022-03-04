package logic

import (
	"golangstudy/jike/awesomeProject/dao/redis"
	"golangstudy/jike/awesomeProject/models"
	"strconv"
)

//投一票加432分 200张赞成票就给帖子续一天
//为帖子投票的函数
//投票的几种情况
/*
d=1  1.之前没有投过票  现在投赞成票 2.之前投反对票，现在投赞成票
d=0  1.之前投赞成票  现在取消  2.之前投反对票 现在取消
d=-1 1.之前没投票  现在投反对票 2.之前投赞成票  现在改成反对票
投票的限制
只有前一个星期允许投票 帖子发布一个星期之后不允许投票
1.到期之后将redis的赞成票数和反对票数存到mysql
2.到期之后删除key
*/

func VoteForPost(userID int64, p *models.ParamVoteData) error {
	// 1. 判断投票限制
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
	// 2.更新分数

	// 3.记录用户为该帖子投票的数据

}
