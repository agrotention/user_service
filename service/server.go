package service

import (
	"github.com/agrotention/user_proto"
	"gorm.io/gorm"
)

type server struct {
	user_proto.UnimplementedUserServiceServer
	db *gorm.DB
}

func NewServer(db *gorm.DB) *server {
	return &server{db: db}
}
