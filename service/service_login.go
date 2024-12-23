package service

import (
	"context"

	"github.com/agrotention/user_proto"
	"github.com/agrotention/user_service/db"
	"github.com/agrotention/user_service/helper"
	"golang.org/x/crypto/bcrypt"
)

func (s *server) UserLogin(
	ctx context.Context,
	req *user_proto.InUserLogin,
) (*user_proto.OutUserLogin, error) {
	if count, err := s.countUsername(req.Username); err != nil {
		return nil, err
	} else if count == 0 {
		return nil, helper.NewServiceError(
			401,
			"invalid username or password",
			"invalid username or password",
		)
	}
	var hash string
	if err := s.db.Model(&db.User{}).Select("password").Find(&hash).Error; err != nil {
		return nil, helper.NewServiceError(500, "internal error", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.GetPassword())); err != nil {
		return nil, helper.NewServiceError(
			401,
			"invalid username or password",
			"invalid username or password",
		)
	}

	return &user_proto.OutUserLogin{AccessToken: "dummyaccesstoken"}, nil
}
