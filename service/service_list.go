package service

import (
	"context"

	"github.com/agrotention/user_proto"
)

func (s *server) UserList(
	ctx context.Context,
	req *user_proto.InUserList,
) (*user_proto.OutUserList, error) {

	return nil, nil
}
