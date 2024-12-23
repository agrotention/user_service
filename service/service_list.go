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
	if req != nil {
		if err := s.db.Limit(int(req.Take)).Offset(int(req.Start)).Find(&users); err != nil {
			log.Println(err.Error)
			return nil, errors.InternalError
		}
	} else {
		if err := s.db.Find(&users); err != nil {
			log.Println(err.Error)
			return nil, errors.InternalError
		}
	}
	resUsers := make([]*user_proto.OutUserDetail, len(users))
	for i, u := range users {
		resUsers[i].Id = u.Id
		resUsers[i].Username = u.Username
		resUsers[i].FullName = u.FullName
	}

	return &user_proto.OutUserList{Users: resUsers}, nil
}
