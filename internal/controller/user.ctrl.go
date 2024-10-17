package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kitae0522/online-library-be/internal/middleware"
	"github.com/kitae0522/online-library-be/internal/model"
	"github.com/kitae0522/online-library-be/internal/repository"
	"github.com/kitae0522/online-library-be/internal/service"
	"github.com/kitae0522/online-library-be/pkg/domain"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{userService: service}
}

func initUserDI(dbconn *model.PrismaClient) *UserController {
	repository := repository.NewUserRepository(dbconn)
	service := service.NewUserService(repository)
	handler := NewUserController(service)
	return handler
}

func initUserRouter(router fiber.Router, handler *UserController) {
	authRouter := router.Group("/user")
	handler.Accessible(authRouter)
	handler.Restricted(authRouter)
}

func (c *UserController) Accessible(router fiber.Router) {
	router.Get("/:userTag", c.GetUserProfile)
}

func (c *UserController) Restricted(router fiber.Router) {
	router.Use(middleware.JWTMiddleware)
	router.Post("/", c.CreateUserProfile)
	router.Patch("/:userTag", c.UpdateUserProfile)
}

func (c *UserController) GetUserProfile(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(domain.DefaultRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		Message:    "✅ 유저 프로필 조회 완료",
	})
}

func (c *UserController) CreateUserProfile(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(domain.DefaultRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		Message:    "✅ 유저 프로필 생성 완료",
	})
}

func (c *UserController) UpdateUserProfile(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(domain.DefaultRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		Message:    "✅ 유저 프로필 수정 완료",
	})
}
