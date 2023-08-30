package repository

import (
	"github.com/Muelsyse/Douyin_Lite/model"
	"gorm.io/gorm"
)

type (
	// UserRepo 用户仓储契约定义
	UserRepo interface {
		QueryUserDuplicate(username string) bool
		CreateUser(name string, password string) (int64, error)
		QueryUserPassword(name string) (int64, string, error)
		QueryUserInfo(userID int64) (*model.User, error)
	}

	userRepo struct{}
)

// NewUserRepo 创建新的仓储实例
func NewUserRepo() UserRepo {
	return new(userRepo)
}

// QueryUserPassword 根据用户名，查询用户ID和密码并返回
func (*userRepo) QueryUserPassword(name string) (int64, string, error) {
	var userSQL model.User
	err := db.Model(&model.User{}).
		Select([]string{"id", "password"}).
		Where("name = ?", name).
		First(&userSQL).Error

	// 发生错误，返回空
	if err != nil {
		return 0, "", err
	}
	return userSQL.ID, userSQL.Password, nil
}

// QueryUserInfo 查询用户, 返回用户
func (*userRepo) QueryUserInfo(userID int64) (*model.User, error) {
	var userSQL model.User
	err := db.Model(&model.User{}).
		Select([]string{"id", "name", "signature", "follow_count", "follower_count", "favorite_count", "total_favorited", "work_count"}).
		Where("id = ?", userID).
		First(&userSQL).Error
	// 发生错误，返回空
	if err != nil {
		return nil, err
	}
	return &userSQL, nil
}

// CreateUser 创建用户, 返回用户ID
func (*userRepo) CreateUser(name string, password string) (int64, error) {
	user := model.User{
		Name:     name,
		Password: password,
	}
	err := db.Model(&model.User{}).Create(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

// QueryUserDuplicate 查询是否用户名重复
func (*userRepo) QueryUserDuplicate(username string) bool {
	var userSQL model.User
	err := db.Model(&model.User{}).
		Where("name = ?", username).
		First(&userSQL).Error
	if err == gorm.ErrRecordNotFound {
		return true
	} else {
		return false
	}

}
