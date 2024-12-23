package service

import (
	"context"

	"github.com/agrotention/user_proto"
	"github.com/agrotention/user_service/db"
	"github.com/agrotention/user_service/helper"
)

func (s *server) UserUpdate(
	ctx context.Context,
	req *user_proto.InUserUpdate,
) (*user_proto.OutUserUpdate, error) {
	// User from request
	user := db.User{
		Id:       req.GetId(),
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		FullName: req.GetFullName(),
	}

	// Check is user exist
	if count, err := s.countId(user.Id); err != nil {
		return nil, err
	} else if count == 0 {
		return nil, helper.NewServiceError(404, "user id not found", "user id not found")
	}

	// Check unique user
	if count, err := s.countUsername(req.Username); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, helper.NewServiceError(409, "username already exist", "username already exist")
	}

	// Query
	if err := s.db.Model(&db.User{}).Where("id = 1", user.Id).Updates(user).Error; err != nil {
		return nil, helper.NewServiceError(500, "internal error", "internal error")
	}

	return &user_proto.OutUserUpdate{
		Id: user.Id,
	}, nil
}
