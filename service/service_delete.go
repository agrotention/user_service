package service

import (
	"context"

	"github.com/agrotention/user_proto"
	"github.com/agrotention/user_service/db"
	"github.com/agrotention/user_service/helper"
)

func (s *server) UserDelete(
	ctx context.Context,
	req *user_proto.InUserDelete,
) (*user_proto.OutUserDelete, error) {
	// Check user exist by helper
	if count, err := s.countId(req.GetId()); err != nil {
		return nil, err
	} else if count == 0 {
		return nil, helper.NewServiceError(
			404,
			"user id not found",
			"user id not found",
		)
	}

	// Query
	if err := s.db.Where("id = ?", req.GetId()).Delete(&db.User{}).Error; err != nil {
		return nil, helper.NewServiceError(500, "internal error", err)
	}

	return &user_proto.OutUserDelete{
		Id: req.GetId(),
	}, nil
}
