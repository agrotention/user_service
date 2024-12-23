package service

import (
	"regexp"

	"github.com/agrotention/user_service/db"
	"github.com/agrotention/user_service/errors"
)

// errors function to validate username
func (s *server) validateUsername(username string) error {
	if len(username) < 3 {
		return errors.NewServiceError(
			400,
			"username minimum length is 3 characters",
			"username minimum length is 3 characters",
		)
	}
	if len(username) > 128 {
		return errors.NewServiceError(
			400,
			"username maximum length is 128 characters",
			"username maximum length is 128 characters",
		)
	}
	return nil
}

func (s *server) validatePassword(password string) error {
	// Check minimum length
	if len(password) < 8 {
		return errors.NewServiceError(
			400,
			"password minimum length is 8 characters",
			"password minimum length is 8 characters",
		)
	}

	// Check for at least one uppercase letter
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	if !uppercaseRegex.MatchString(password) {
		return errors.NewServiceError(
			400,
			"password must contain at least one uppercase letter",
			"password must contain at least one uppercase letter",
		)
	}

	// Check for at least one number
	numberRegex := regexp.MustCompile(`[0-9]`)
	if !numberRegex.MatchString(password) {
		return errors.NewServiceError(
			400,
			"password must contain at least one number",
			"password must contain at least one number",
		)
	}

	return nil
}

// errors function to check if username already exists
func (s *server) countUsername(username string) (int64, error) {
	var count int64
	if err := s.db.Model(&db.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return 0, errors.NewServiceError(500, "internal error", err)
	}
	return count, nil
}

// errors function to check if username already exists
func (s *server) countId(id string) (int64, error) {
	var count int64
	if err := s.db.Model(&db.User{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return 0, errors.NewServiceError(500, "internal error", err)
	}
	return count, nil
}
