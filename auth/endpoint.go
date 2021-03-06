package auth

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type registerRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type registerResponse struct {
	Err error `json:"error,omitempty"`
}

func (res *registerResponse) error() error { return res.Err }

func makeRegisterEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(registerRequest)
		err := s.Register(ctx, req.Name, req.Password, req.Email)
		return registerResponse{Err: err}, nil
	}
}

type unregisterRequest struct {
	UID      uint64 `json:"uid"`
	Password string `json:"password"`
	SID      string `json:"sid"`
}

type unregisterResponse struct {
	Err error `json:"error,omitempty"`
}

func (res *unregisterResponse) error() error { return res.Err }

func makeUnregisterEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(unregisterRequest)
		err := s.Unregister(ctx, req.UID, req.Password, req.SID)
		return unregisterResponse{Err: err}, nil
	}
}

type loginRequest struct {
	NameOrEmail string `json:"name_or_email"`
	Password    string `json:"password"`
}

type loginResponse struct {
	Err error `json:"error,omitempty"`
}

func (res *loginResponse) error() error { return res.Err }

func makeLoginEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loginRequest)
		sess, err := s.Login(ctx, req.NameOrEmail, req.Password)
		return loginResponse{Err: err}, nil
	}
}

type logoutRequest struct {
	UID uint64 `json:"uid"`
	SID string `json:"sid"`
}

type logoutResponse struct {
	Err error `json:"error,omitempty"`
}

func (res *logoutResponse) error() error { return res.Err }

func makeLogoutEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(logoutRequest)
		err := s.Logout(ctx, req.UID, req.SID)
		if err != nil {
			return logoutResponse{Err: err}, nil
		}
	}
}
