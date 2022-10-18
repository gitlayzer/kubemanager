package service

import (
	"errors"
	"kubemanager/LoadFiles"

	"github.com/wonderivan/logger"
)

var Login login

type login struct{}

func (l *login) Auth(username, password string) (err error) {
	if username == LoadFiles.ReadAdmin() && password == LoadFiles.ReadPassword() {
		return nil
	} else {
		logger.Error("登录失败, 用户名或密码错误")
		return errors.New("登录失败, 用户名或密码错误")
	}
	return nil
}
