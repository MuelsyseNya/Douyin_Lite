package repository

import (
	"github.com/Muelsyse/Douyin_Lite/model"
	"gorm.io/gorm"
)

type (
	// UserRepo 用户仓储契约定义
	UserRepo interface {
		QueryUserDuplicate(string) bool
		CreateUser(*model.User) (int64, error)
		QueryUserPassword(string) (int64, string, error)
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
func (*userRepo) QueryUserInfo(fields []string, conditions *model.User) (*model.User, error) {
	var userSQL model.User
	err := db.Model(&model.User{}).
		Select(fields).
		Where(conditions).
		First(&userSQL).Error

	// 发生错误，返回空
	if err != nil {
		return nil, err
	}
	return &userSQL, nil
}

// CreateUser 创建用户, 返回用户ID
func (*userRepo) CreateUser(user *model.User) (int64, error) {
	err := db.Model(&model.User{}).Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

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
