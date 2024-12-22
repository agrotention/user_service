package service

import (
	"context"

	"github.com/agrotention/user_proto"
	"github.com/agrotention/user_service/db"
)

func (s *server) UserDelete(
	ctx context.Context,
	req *user_proto.InUserDelete,
) (*user_proto.OutUserDelete, error) {

	return nil, nil
}

func (s *server) UserDetail(
	ctx context.Context,
	req *user_proto.InUserDetail,
) (*user_proto.OutUserDetail, error) {

	return nil, nil
}

func (s *server) UserDisable(
	ctx context.Context,
	req *user_proto.InUserDisable,
) (*user_proto.OutUserDisable, error) {

	return nil, nil
}

func (s *server) UserList(
	ctx context.Context,
	req *user_proto.InUserList,
) (*user_proto.OutUserList, error) {

	return nil, nil
}

func (s *server) UserLogin(
	ctx context.Context,
	req *user_proto.InUserLogin,
) (*user_proto.OutUserLogin, error) {

	return nil, nil
}

func (s *server) UserRegister(
	ctx context.Context,
	req *user_proto.InUserRegister,
) (*user_proto.OutUserRegister, error) {
	// Validate Username
	if err := s.validateUsername(req.Username); err != nil {
		return nil, err
	}

	// Validate Password
	if err := s.validatePassword(req.Password); err != nil {
		return nil, err
	}

	// Check if username exists
	if err := s.checkUsernameExists(req.Username); err != nil {
		return nil, err
	}

	// Create the user
	newUser := &db.User{
		Username: req.Username,
		Password: req.Password,
		FullName: req.FullName,
	}
	if err := s.db.Create(newUser).Error; err != nil {
		return nil, err
	}

	// Return successful response
	return &user_proto.OutUserRegister{
		Id: newUser.Id,
	}, nil
}

func (s *server) UserUpdate(
	ctx context.Context,
	req *user_proto.InUserUpdate,
) (*user_proto.OutUserUpdate, error) {

	return nil, nil
}
