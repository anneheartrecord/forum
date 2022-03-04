package logic

import (
	"golangstudy/jike/awesomeProject/dao/mysql"
	"golangstudy/jike/awesomeProject/models"
)

func GetCommunityList() ([]*models.Community, error) {
	//查找数据库  查到所有的community 并返回
	return mysql.GetCommunityList()

}
func GetCommunityDetail(id int64) (data *models.CommunityDetail, err error) {
	return mysql.GetCommunityDetailByID(id)
}
