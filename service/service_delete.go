package service

import (
	"context"

	"github.com/agrotention/user_proto"
)

func (s *server) UserDelete(
	ctx context.Context,
	req *user_proto.InUserDelete,
) (*user_proto.OutUserDelete, error) {
	// TODO Create UserDelete service
	panic("unimplemented!")
}
