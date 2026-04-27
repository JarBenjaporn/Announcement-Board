package validators

import (
	"announcement-board/models"
	"errors"
	"strings"
)

func ValidateCreateRequest(req *models.CreateRequest) error {
	if 	strings.TrimSpace(req.Title) == "" ||
		strings.TrimSpace(req.Body) == "" ||
		strings.TrimSpace(req.Author) == "" {
		return errors.New("title, body, and author are required cannot be empty")
	}
	return nil
}