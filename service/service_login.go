package service

import (
	"context"
	"log"

	"github.com/agrotention/user_proto"
	"github.com/agrotention/user_service/auth"
	"github.com/agrotention/user_service/db"
	"github.com/agrotention/user_service/helper/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *server) UserLogin(
	ctx context.Context,
	req *user_proto.InUserLogin,
) (*user_proto.OutUserLogin, error) {
	InvalidUsernameOrPassword := errors.NewServiceError(401, "invalid username or password")

	if count, err := s.countUsername(req.Username); err != nil {
		log.Println(err.Error())
		return nil, err
	} else if count == 0 {
		return nil, InvalidUsernameOrPassword
	}

	// Query user
	var user db.User
	if err := s.db.Where("username = ?", req.GetUsername()).Find(&user).Error; err != nil {
		log.Println(err.Error())
		return nil, errors.InternalError
	}

	// Verify Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword())); err != nil {
		log.Println(err.Error())
		return nil, InvalidUsernameOrPassword
	}

	// Generate token
	claims := auth.NewClaims(user.Id)
	token, err := auth.CreateToken(claims)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.InternalError
	}

	return &user_proto.OutUserLogin{AccessToken: token}, nil
}
