package service

import (
	"context"

	"github.com/agrotention/user_proto"
	"github.com/agrotention/user_service/db"
	"github.com/agrotention/user_service/helper/errors"
)

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
	if count, err := s.countUsername(req.Username); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.NewServiceError(409, "username already exist", "username already exist")
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
