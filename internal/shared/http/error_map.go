package http

import (
	"errors"
	custom_errors "hotel_system2/internal/shared/errors"

	"github.com/gofiber/fiber/v2"
)

func StatusFor(err error) int {
	var customErr *custom_errors.ErrorResponse

	if errors.As(err, &customErr) {
		return customErr.StatusCode
	}

	return fiber.StatusInternalServerError
}