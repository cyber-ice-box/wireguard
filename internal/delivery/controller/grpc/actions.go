package grpc

import (
	"context"
<<<<<<< HEAD
	"github.com/cyber-ice-box/wireguard/pkg/controller/grpc/protobuf"
=======
	"gitlab.com/cyber-ice-box/wireguard/pkg/controller/grpc/protobuf"
>>>>>>> origin/develop
)

type service interface {
	AddClient(name, ip, publicKey string) error
	DeleteClient(name string) error
	AllowClientAccess(name, destCIDR string) error
	DenyClientAccess(name string) error
}

func (w *Wireguard) AddUser(_ context.Context, request *protobuf.AddUserRequest) (*protobuf.EmptyResponse, error) {
	if err := w.service.AddClient(request.Name, request.Ip, request.PublicKey); err != nil {
		return &protobuf.EmptyResponse{}, err
	}

	return &protobuf.EmptyResponse{}, nil
}

func (w *Wireguard) DeleteUser(_ context.Context, request *protobuf.DeleteUserRequest) (*protobuf.EmptyResponse, error) {
	if err := w.service.DeleteClient(request.Name); err != nil {
		return &protobuf.EmptyResponse{}, err
	}

	return &protobuf.EmptyResponse{}, nil
}

func (w *Wireguard) AllowAccess(_ context.Context, request *protobuf.UserAllowAccessRequest) (*protobuf.EmptyResponse, error) {
	if err := w.service.AllowClientAccess(request.Name, request.DestCIDR); err != nil {
		return &protobuf.EmptyResponse{}, err
	}
	return &protobuf.EmptyResponse{}, nil
}

func (w *Wireguard) DenyAccess(_ context.Context, request *protobuf.UserDenyAccessRequest) (*protobuf.EmptyResponse, error) {
	if err := w.service.DenyClientAccess(request.Name); err != nil {
		return &protobuf.EmptyResponse{}, err
	}
	return &protobuf.EmptyResponse{}, nil
}
