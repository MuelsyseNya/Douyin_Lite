package service

import (
	"github.com/Muelsyse/Douyin_Lite/middleware"
	"github.com/Muelsyse/Douyin_Lite/repository"
	"github.com/Muelsyse/Douyin_Lite/response"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func NewUser() *UserService {
	return new(UserService)
}

// lengthCheck 用户名和密码长度检查
func (*UserService) lengthCheck(name string, password string) bool {
	if len(name) == 0 || len(name) > 24 || len(password) == 0 || len(password) > 24 {
		return false
	}
	return true
}

func (u *UserService) Register(name string, password string) *response.UserLoginResponse {
	// 长度检查
	if ok := u.lengthCheck(name, password); !ok {
		return &response.UserLoginResponse{
			Response: response.ErrPasswordLength,
		}
	}

	if ok := repository.NewUserRepo().QueryUserDuplicate(name); ok {
		// 用户名无重复
		// 密码加密
		passwordByte, err := bcrypt.GenerateFromPassword([]byte(password), 8)
		if err != nil {
			return &response.UserLoginResponse{
				Response: response.ErrPasswordEncrypt,
			}
		}
		password = string(passwordByte)
	} else {
		return &response.UserLoginResponse{
			Response: response.ErrDuplicatedName,
		}
	}

	// 数据库中插入用户
	userID, err := repository.NewUserRepo().CreateUser(name, password)
	if err != nil {
		return &response.UserLoginResponse{
			Response: response.ErrUserCreation,
		}
	}

	// 设置token
	token, err := middleware.GenerateToken(userID)
	if err != nil {
		return &response.UserLoginResponse{
			Response: response.ErrTokenGenerate,
		}
	}

	// 成功后，返回响应
	return &response.UserLoginResponse{
		Response: response.OK,
		UserId:   &userID,
		Token:    &token,
	}
}

func (u *UserService) Login(name string, password string) *response.UserLoginResponse {
	// 长度检查
	if ok := u.lengthCheck(name, password); !ok {
		return &response.UserLoginResponse{
			Response: response.ErrPasswordLength,
		}
	}

	// 从数据库中查询用户名对应ID和密码
	userId, userPassword, err := repository.NewUserRepo().QueryUserPassword(name)
	if err != nil {
		return &response.UserLoginResponse{
			Response: response.ErrLoginQuery,
		}
	}

	// 和数据库中的加密密码进行比较
	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	if err != nil {
		return &response.UserLoginResponse{
			Response: response.ErrPassword,
		}
	}

	// 设置token
	token, err := middleware.GenerateToken(userId)
	if err != nil {
		return &response.UserLoginResponse{
			Response: response.ErrTokenGenerate,
		}
	}

	// 成功后，返回响应
	return &response.UserLoginResponse{
		Response: response.OK,
		UserId:   &userId,
		Token:    &token,
	}
}

func (*UserService) QueryUserInfo(toUserID int64, UserID int64) *response.UserResponse {
	user, err := repository.NewUserRepo().QueryUserInfo(toUserID)
	if err != nil {
		return &response.UserResponse{
			Response: response.ErrQueryUserInfo,
		}
	}
	isFollow, err := repository.NewRelationRepo().IsFollow(UserID, toUserID)
	if err != nil {
		return &response.UserResponse{
			Response: response.ErrQueryIfFollow,
		}
	}
	// 查询完毕，返回响应
	return &response.UserResponse{
		Response: response.OK,
		User: &response.User{
			Avatar:          "",
			BackgroundImage: "",
			FavoriteCount:   user.FavoriteCount,
			FollowerCount:   user.FollowerCount,
			FollowCount:     user.FollowCount,
			ID:              UserID,
			IsFollow:        isFollow,
			Name:            user.Name,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
		},
	}
}
