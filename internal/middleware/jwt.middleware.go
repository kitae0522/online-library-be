package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/kitae0522/online-library-be/pkg/crypt"
	"github.com/kitae0522/online-library-be/pkg/utils"
)

func JWTMiddleware(ctx *fiber.Ctx) error {
	authHeader := strings.Split(ctx.Get("Authorization"), " ")
	if len(authHeader) != 2 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	token := authHeader[1]

	uuid, err := crypt.ParseJWT(token)
	if err != nil {
		return utils.CreateErrorRes(ctx, fiber.StatusUnauthorized, "❌ 유효하지 않는 토큰 값입니다.", err)
	}
	ctx.Locals("uuid", uuid)
	return ctx.Next()
}

func GetUUIDFromMiddleware(ctx *fiber.Ctx) string {
	return ctx.Locals("uuid").(string)
}
