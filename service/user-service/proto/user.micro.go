// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user.proto

/*
Package shop_srv_user is a generated protocol buffer package.

It is generated from these files:
	proto/user.proto

It has these top-level messages:
	LoginReq
	UserResp
	PhoneCodeReq
	PhoneCodeResp
*/
package shop_srv_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*UserResp, error)
	GetCode(ctx context.Context, in *PhoneCodeReq, opts ...client.CallOption) (*PhoneCodeResp, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "shop.srv.user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Login(ctx context.Context, in *LoginReq, opts ...client.CallOption) (*UserResp, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(UserResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetCode(ctx context.Context, in *PhoneCodeReq, opts ...client.CallOption) (*PhoneCodeResp, error) {
	req := c.c.NewRequest(c.name, "User.GetCode", in)
	out := new(PhoneCodeResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Login(context.Context, *LoginReq, *UserResp) error
	GetCode(context.Context, *PhoneCodeReq, *PhoneCodeResp) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Login(ctx context.Context, in *LoginReq, out *UserResp) error
		GetCode(ctx context.Context, in *PhoneCodeReq, out *PhoneCodeResp) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Login(ctx context.Context, in *LoginReq, out *UserResp) error {
	return h.UserHandler.Login(ctx, in, out)
}

func (h *userHandler) GetCode(ctx context.Context, in *PhoneCodeReq, out *PhoneCodeResp) error {
	return h.UserHandler.GetCode(ctx, in, out)
}
