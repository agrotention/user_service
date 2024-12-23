package service

import (
	"context"
	"log"

	"github.com/agrotention/user_proto"
	"github.com/agrotention/user_service/db"
	"github.com/agrotention/user_service/helper/errors"
)

func (s *server) UserDetail(
	ctx context.Context,
	req *user_proto.InUserDetail,
) (*user_proto.OutUserDetail, error) {

	// Check user exist by errors
	if count, err := s.countId(req.GetId()); err != nil {
		return nil, err
	} else if count == 0 {
		return nil, errors.UserNotFound
	}

	// Query
	var user db.User
	if err := s.db.First(&user, "id = ?", req.GetId()).Error; err != nil {
		log.Println(err)
		return nil, errors.InternalError
	}

	return &user_proto.OutUserDetail{
		Id:       user.Id,
		Username: user.Username,
		FullName: user.FullName,
	}, nil
}
