package utils

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorRes struct {
	IsError    bool        `json:"isError"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Error      interface{} `json:"error"`
}

func CreateErrorRes(ctx *fiber.Ctx, statusCode int, errMessage string, err interface{}) error {
	return ctx.Status(statusCode).JSON(ErrorRes{
		IsError:    true,
		StatusCode: statusCode,
		Message:    errMessage,
		Error:      err,
	})
}
