package service

import (
	"context"
	"log"

	"github.com/agrotention/user_proto"
	"github.com/agrotention/user_service/db"
	"github.com/agrotention/user_service/helper/errors"
)

func (s *server) UserDisable(
	ctx context.Context,
	req *user_proto.InUserDisable,
) (*user_proto.OutUserDisable, error) {
	// Check user exist by errors
	if count, err := s.countId(req.GetId()); err != nil {
		return nil, err
	} else if count == 0 {
		return nil, errors.UserNotFound
	}

	// Query
	if err := s.db.Where("id = ?", req.GetId()).Delete(&db.User{}).Error; err != nil {
		log.Println(err.Error())
		return nil, errors.InternalError
	}

	return &user_proto.OutUserDisable{
		Id: req.GetId(),
	}, nil
}
