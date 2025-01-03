package service

import (
	"context"

	"github.com/agrotention/user_proto"
	"github.com/agrotention/user_service/db"
	"github.com/agrotention/user_service/helper/errors"
)

func (s *server) UserDelete(
	ctx context.Context,
	req *user_proto.InUserDelete,
) (*user_proto.OutUserDelete, error) {
	// Check user exist by service helper
	if count, err := s.countId(req.GetId()); err != nil {
		return nil, err
	} else if count == 0 {
		return nil, errors.UserNotFound
	}

	// Query
	if err := s.db.Unscoped().Where("id = ?", req.GetId()).Delete(&db.User{}).Error; err != nil {
		return nil, errors.InternalError
	}

	return &user_proto.OutUserDelete{
		Id: req.GetId(),
	}, nil
}
