package service

import (
	"errors"
	"github.com/Muelsyse/Douyin_Lite/middleware"
	"github.com/Muelsyse/Douyin_Lite/model"
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

func (u *UserService) Register(name string, password string) (*response.UserLoginResponse, error) {
	// 长度检查
	if ok := u.lengthCheck(name, password); !ok {
		return &response.UserLoginResponse{
			Response: response.ErrPasswordLength,
		}, errors.New("account name or password length is illegal")
	}

	if ok := repository.NewUserRepo().QueryUserDuplicate(name); ok {
		// 用户名无重复
		// 密码加密
		passwordByte, err := bcrypt.GenerateFromPassword([]byte(password), 8)
		if err != nil {
			return &response.UserLoginResponse{
				Response: response.ErrPasswordEncrypt,
			}, err
		}
		password = string(passwordByte)
	} else {
		return &response.UserLoginResponse{
			Response: response.ErrDuplicatedName,
		}, errors.New("user name already exists")
	}

	// 数据库中插入用户
	userID, err := repository.NewUserRepo().CreateUser(&model.User{
		Name:     name,
		Password: password,
	})
	if err != nil {
		return &response.UserLoginResponse{
			Response: response.ErrUserCreation,
		}, err
	}

	// 设置token
	token, err := middleware.GenerateToken(userID)
	if err != nil {
		return &response.UserLoginResponse{
			Response: response.ErrTokenGenerate,
		}, err
	}

	// 成功后，返回响应
	return &response.UserLoginResponse{
		Response: response.OK,
		UserId:   &userID,
		Token:    &token,
	}, err
}

func (u *UserService) Login(name string, password string) (*response.UserLoginResponse, error) {
	// 长度检查
	if ok := u.lengthCheck(name, password); !ok {
		return &response.UserLoginResponse{
			Response: response.ErrPasswordLength,
		}, errors.New("account name or password length is illegal")
	}

	// 从数据库中查询用户名对应ID和密码
	userId, userPassword, err := repository.NewUserRepo().QueryUserPassword(name)
	if err != nil {
		return &response.UserLoginResponse{
			Response: response.ErrLoginQuery,
		}, err
	}

	// 和数据库中的加密密码进行比较
	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	if err != nil {
		return &response.UserLoginResponse{
			Response: response.ErrPassword,
		}, err
	}

	// 设置token
	token, err := middleware.GenerateToken(userId)
	if err != nil {
		return &response.UserLoginResponse{
			Response: response.ErrTokenGenerate,
		}, err
	}

	// 成功后，返回响应
	return &response.UserLoginResponse{
		Response: response.OK,
		UserId:   &userId,
		Token:    &token,
	}, err
}
