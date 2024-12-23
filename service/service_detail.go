package service

import (
	"context"

	"github.com/agrotention/user_proto"
	"github.com/agrotention/user_service/db"
	"github.com/agrotention/user_service/errors"
)

func (s *server) UserDetail(
	ctx context.Context,
	req *user_proto.InUserDetail,
) (*user_proto.OutUserDetail, error) {

	// Check user exist by errors
	if count, err := s.countId(req.GetId()); err != nil {
		return nil, err
	} else if count == 0 {
		return nil, errors.NewServiceError(
			404,
			"user id not found",
			"user id not found",
		)
	}

	// Query
	var user db.User
	if err := s.db.First(&user, "id = ?", req.GetId()).Error; err != nil {
		return nil, errors.NewServiceError(500, "internal error", err)
	}

	return &user_proto.OutUserDetail{
		Id:       user.Id,
		Username: user.Username,
		FullName: user.FullName,
	}, nil
}
