package handler

import (
	"context"
	"github.com/micro/go-micro/errors"
	"go-micro-shopping/user/model"
	user "go-micro-shopping/user/proto/user"
	"go-micro-shopping/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Repo *repository.User
}

func (h *User) Register(ctx context.Context, in *user.RegisterRequest, out *user.Response) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(in.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &model.User{
		Name:     in.User.Name,
		Phone:    in.User.Phone,
		Password: string(hashPassword),
	}
	if err := h.Repo.Create(user); err != nil {
		return nil
	}
	out.Code = "200"
	out.Msg = "注册成功"

	return nil
}

func (h *User) Login(ctx context.Context, in *user.LoginRequest, out *user.Response) error {
	user, err := h.Repo.FindByField("phone", in.Phone, "id , password")
	if err != nil {
		return nil
	}
	if user == nil {
		return errors.Unauthorized("go.micro.srv.user.login", "该手机号不存在")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return errors.Unauthorized("go.micro.srv.user.login", "密码错误")
	}
	out.Code = "200"
	out.Msg = "登录成功"

	return nil
}

func (h *User) UpdatePassword(ctx context.Context, in *user.UpdatePasswordRequest, out *user.Response) error {
	user, err := h.Repo.Find(in.Uid)
	if err != nil {
		return errors.Unauthorized("go.micro.srv.user.login", "没有该用户")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.OldPassword)); err != nil {
		return errors.Unauthorized("go.micro.srv.user.login", "旧密码验证错误")
	}
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPwd)
	h.Repo.Update(user)

	out.Code = "200"
	out.Msg = user.Name + "，您的密码更新成功"

	return nil
}
