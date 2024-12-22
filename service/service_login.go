package service

import (
	"context"

	"github.com/agrotention/user_proto"
)

func (s *server) UserLogin(
	ctx context.Context,
	req *user_proto.InUserLogin,
) (*user_proto.OutUserLogin, error) {
	// TODO Create UserLogin service
	panic("unimplemented!")
}
