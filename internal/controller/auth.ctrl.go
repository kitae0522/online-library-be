package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/kitae0522/online-library-be/internal/dto"
	"github.com/kitae0522/online-library-be/internal/middleware"
	"github.com/kitae0522/online-library-be/internal/model"
	"github.com/kitae0522/online-library-be/internal/repository"
	"github.com/kitae0522/online-library-be/internal/service"
	"github.com/kitae0522/online-library-be/pkg/utils"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{authService: service}
}

func initAuthDI(dbconn *model.PrismaClient) *AuthController {
	repository := repository.NewAuthRepository(dbconn)
	service := service.NewAuthService(repository)
	handler := NewAuthController(service)
	return handler
}

func initAuthRouter(router fiber.Router, handler *AuthController) {
	authRouter := router.Group("/auth")
	handler.Accessible(authRouter)
	handler.Restricted(authRouter)
}

func (c *AuthController) Accessible(router fiber.Router) {
	router.Post("/register", c.Register)
	router.Post("/login", c.Login)
}

func (c *AuthController) Restricted(router fiber.Router) {
	router.Use(middleware.JWTMiddleware)
	router.Get("/test", c.Login)
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	user := new(dto.AuthRegisterReq)
	if err := utils.Bind(ctx, user); err != nil {
		return utils.CreateErrorRes(ctx, fiber.StatusBadRequest, "❌ 회원가입 실패. Body Binding 과정에서 문제 발생", err)
	}

	err := c.authService.Register(user)
	if err != nil {
		if _, uniqueErr := model.IsErrUniqueConstraint(err); uniqueErr {
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 회원가입 실패. 중복된 유저가 존재합니다.", err)
		}
		return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 회원가입 실패. Repository에서 문제 발생", err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.AuthRegisterRes{
		IsError:    false,
		StatusCode: fiber.StatusCreated,
		Message:    "✅ 회원가입 완료",
	})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	user := new(dto.AuthLoginReq)
	if err := utils.Bind(ctx, user); err != nil {
		return utils.CreateErrorRes(ctx, fiber.StatusBadRequest, "❌ 로그인 실패. Body Binding 과정에서 문제 발생", err)
	}

	token, err := c.authService.Login(user)
	if err != nil {
		log.Printf("%v", err)
		switch err {
		case model.ErrNotFound:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 로그인 실패. 존재하지 않는 사용자입니다.", err)
		case utils.ErrWrongPassword:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 로그인 실패. 패스워드가 일치하지 않습니다.", err)
		default:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 로그인 실패. Repository에서 문제 발생", err)
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.AuthLoginRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		Message:    "✅ 로그인 완료",
		Token:      token,
	})
}
