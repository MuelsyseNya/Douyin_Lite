package repository

import (
	"github.com/Muelsyse/Douyin_Lite/model"
	"gorm.io/gorm"
)

type (
	// RelationRepo 用户仓储契约定义
	RelationRepo interface {
		IsFollow(userID int64, toUserID int64) (bool, error)
	}

	relationRepo struct{}
)

// NewRelationRepo 创建新的仓储实例
func NewRelationRepo() RelationRepo {
	return new(relationRepo)
}

func (*relationRepo) IsFollow(userID int64, toUserID int64) (bool, error) {
	var followSQL model.Follow
	err := db.Model(&model.Follow{}).
		Where("user_id = ? and to_user_id = ?", userID, toUserID).
		First(&followSQL).Error
	if err == gorm.ErrRecordNotFound {
		// 没有找到记录，说明没有关注
		return false, nil
	} else if err == nil {
		// 没有发生错误，说明查到关注记录
		return true, nil
	} else {
		// 发生意外错误
		return false, err
	}
}
