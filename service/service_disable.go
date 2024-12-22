package service

import (
	"context"

	"github.com/agrotention/user_proto"
)

func (s *server) UserDisable(
	ctx context.Context,
	req *user_proto.InUserDisable,
) (*user_proto.OutUserDisable, error) {
	// TODO Create UserDisable service
	panic("unimplemented!")
}
