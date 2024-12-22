package service

import (
	"context"

	"github.com/agrotention/user_proto"
)

func (s *server) UserUpdate(
	ctx context.Context,
	req *user_proto.InUserUpdate,
) (*user_proto.OutUserUpdate, error) {
	// TODO Create UserUpdate service
	panic("unimplemented!")
}
