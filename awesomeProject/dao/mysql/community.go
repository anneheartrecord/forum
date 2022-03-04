package mysql

import (
	"database/sql"
	"golangstudy/jike/awesomeProject/models"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}
func GetCommunityDetailByID(id int64) (communitydetail *models.CommunityDetail, err error) {
	communitydetail = new(models.CommunityDetail)
	sqlstr := `select community_id,community_name,introduction,create_time from community where community_id=?`
	if err = db.Get(communitydetail, sqlstr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
		return communitydetail, err
	}
	return
}
