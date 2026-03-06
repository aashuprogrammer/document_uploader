package api

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func errorHandler(c fiber.Ctx, err error) error {
	// Default status code
	code := fiber.StatusInternalServerError

	// If Fiber error, override status
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return c.Status(code).JSON(fiber.Map{
		"error":   true,
		"message": err.Error(),
	})
}

func NotFoundError(message string) *fiber.Error {
	return &fiber.Error{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func InternalServerError(message string) *fiber.Error {
	return &fiber.Error{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func BadRequestError(message string) *fiber.Error {
	return &fiber.Error{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func ValidationError(message string) *fiber.Error {
	return &fiber.Error{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}
