package service

import (
	"context"

	"github.com/agrotention/user_proto"
)

func (s *server) UserList(
	ctx context.Context,
	req *user_proto.InUserList,
) (*user_proto.OutUserList, error) {
	// TODO Create UserList service
	panic("unimplemented!")
}
