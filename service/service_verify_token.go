package service

import (
	"context"

	"github.com/agrotention/user_proto"
	"github.com/agrotention/user_service/auth"
	"github.com/agrotention/user_service/helper/errors"
)

func (s *server) UserVerifyToken(
	ctx context.Context,
	req *user_proto.InUserVerifyToken,
) (*user_proto.OutUserVerifyToken, error) {
	if req == nil {
		return nil, errors.NewServiceError(400, "empty message")
	}
	claims, err := auth.ParseToken(req.Token)
	if err != nil {
		return nil, errors.NewServiceError(401, "unauthorized")
	}

	return &user_proto.OutUserVerifyToken{
		Exp: claims.Exp.Unix(),
		Iat: claims.Iat.Unix(),
		Nb:  claims.Nbf.Unix(),
		Iss: claims.Iss,
		Sub: claims.Sub,
		Aud: claims.Aud,
	}, nil
}
