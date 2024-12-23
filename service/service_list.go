package service

import (
	"context"
	"log"

	"github.com/agrotention/user_proto"
	"github.com/agrotention/user_service/db"
	"github.com/agrotention/user_service/helper/errors"
)

func (s *server) UserList(
	ctx context.Context,
	req *user_proto.InUserList,
) (*user_proto.OutUserList, error) {
	var users []db.User
	query := s.db.Model(&db.User{})

	// Apply pagination if request specifies it
	if req != nil {
		if req.Take > 0 {
			query = query.Limit(int(req.Take))
		}
		if req.Start > 0 {
			query = query.Offset(int(req.Start))
		}
	}

	// Execute query
	if err := query.Find(&users).Error; err != nil {
		log.Printf("Error fetching users: %v\n", err)
		return nil, errors.InternalError
	}

	// Transform users to proto format
	resUsers := make([]*user_proto.OutUserDetail, len(users))
	for i, u := range users {
		resUsers[i] = &user_proto.OutUserDetail{
			Id:       u.Id,
			Username: u.Username,
			FullName: u.FullName,
		}
	}

	return &user_proto.OutUserList{Users: resUsers}, nil
}
