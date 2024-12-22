package service

import (
	"context"

	"github.com/agrotention/user_proto"
)

func (s *server) UserDetail(
	ctx context.Context,
	req *user_proto.InUserDetail,
) (*user_proto.OutUserDetail, error) {
	// TODO Create UserDetail service
	panic("unimplemented!")
}
