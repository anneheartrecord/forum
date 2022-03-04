package redis

//redis key

//redis 相关的key 尽量使用命名空间的方式 区分 方便查询和拆分
const (
	Prefix             = "awesomeProject:"
	KeyPostTimeZSet    = "post:time"  //zset 以帖子发帖时间的分数
	KeyPostScoreZSet   = "post:score" //zset 帖子投票的分数
	KeyPostVotedZSetPF = "post:voted" //参数是post id 记录用户及投票类型
)

func getRedisKey(key string) string {
	return Prefix + key
}
