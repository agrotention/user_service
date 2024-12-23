package service

import (
	"context"
	"log"

	"github.com/agrotention/user_proto"
	"github.com/agrotention/user_service/db"
	"github.com/agrotention/user_service/helper/errors"
)

func (s *server) UserUpdate(
	ctx context.Context,
	req *user_proto.InUserUpdate,
) (*user_proto.OutUserUpdate, error) {

	// Check is user exist
	if count, err := s.countId(req.GetId()); err != nil {
		return nil, err
	} else if count == 0 {
		return nil, errors.UserNotFound
	}

	// Validate Username
	if err := s.validateUsername(req.GetUsername()); err != nil {
		return nil, err
	}

	// Validate password
	if err := s.validatePassword(req.GetPassword()); err != nil {
		return nil, err
	}

	// Check unique username
	if count, err := s.countUsername(req.Username); err != nil {
		return nil, err
	} else if count != 0 {
		return nil, errors.NewServiceError(409, "username already exist")
	}

	// Convert User from request
	user := db.User{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		FullName: req.GetFullName(),
	}
	// Query
	if err := s.db.Model(&db.User{}).Where("id = ?", req.GetId()).Updates(user).Error; err != nil {
		log.Println(err.Error())
		return nil, errors.InternalError
	}

	return &user_proto.OutUserUpdate{
		Id: user.Id,
	}, nil
}
