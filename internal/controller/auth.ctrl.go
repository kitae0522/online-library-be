package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/kitae0522/online-library-be/internal/middleware"
	"github.com/kitae0522/online-library-be/internal/model"
	"github.com/kitae0522/online-library-be/internal/repository"
	"github.com/kitae0522/online-library-be/internal/service"
	"github.com/kitae0522/online-library-be/pkg/domain"
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
	router.Patch("/reset", c.PasswordReset)
	router.Delete("/withdraw", c.Withdraw)
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	createUserPayload := new(domain.AuthRegisterReq)
	if err := utils.Bind(ctx, createUserPayload); err != nil {
		return utils.CreateErrorRes(ctx, fiber.StatusBadRequest, "❌ 회원가입 실패. Body Binding 과정에서 문제 발생", err)
	}

	err := c.authService.Register(createUserPayload)
	if err != nil {
		if _, uniqueErr := model.IsErrUniqueConstraint(err); uniqueErr {
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 회원가입 실패. 중복된 유저가 존재합니다.", err)
		}
		return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 회원가입 실패. Repository에서 문제 발생", err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(domain.DefaultRes{
		IsError:    false,
		StatusCode: fiber.StatusCreated,
		Message:    "✅ 회원가입 완료",
	})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	loginUserPayload := new(domain.AuthLoginReq)
	if err := utils.Bind(ctx, loginUserPayload); err != nil {
		return utils.CreateErrorRes(ctx, fiber.StatusBadRequest, "❌ 로그인 실패. Body Binding 과정에서 문제 발생", err)
	}

	token, err := c.authService.Login(loginUserPayload)
	if err != nil {
		log.Printf("%v", err)
		switch err {
		case model.ErrNotFound:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 로그인 실패. 존재하지 않는 사용자입니다.", err)
		case domain.ErrWrongPassword:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 로그인 실패. 패스워드가 일치하지 않습니다.", err)
		default:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 로그인 실패. Repository에서 문제 발생", err)
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.AuthLoginRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		Message:    "✅ 로그인 완료",
		Token:      token,
	})
}

func (c *AuthController) PasswordReset(ctx *fiber.Ctx) error {
	resetPayload := new(domain.AuthPasswordResetReq)
	if err := utils.Bind(ctx, resetPayload); err != nil {
		return utils.CreateErrorRes(ctx, fiber.StatusBadRequest, "❌ 비밀번호 초기화 실패. Body Binding 과정에서 문제 발생", err)
	}

	if err := c.authService.PasswordReset(resetPayload); err != nil {
		log.Printf("%v", err)
		switch err {
		case model.ErrNotFound:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 비밀번호 초기화 실패. 존재하지 않는 사용자입니다.", err)
		case domain.ErrWrongPassword:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 비밀번호 초기화 실패. 패스워드가 일치하지 않습니다.", err)
		default:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 비밀번호 초기화 실패. Repository에서 문제 발생", err)
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.DefaultRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		Message:    "✅ 비밀번호 초기화 완료",
	})
}

func (c *AuthController) Withdraw(ctx *fiber.Ctx) error {
	withdrawPayload := new(domain.AuthWithdrawReq)
	if err := utils.Bind(ctx, withdrawPayload); err != nil {
		return utils.CreateErrorRes(ctx, fiber.StatusBadRequest, "❌ 유저 탈퇴 실패. Body Binding 과정에서 문제 발생", err)
	}

	if err := c.authService.Withdraw(withdrawPayload); err != nil {
		log.Printf("%v", err)
		switch err {
		case model.ErrNotFound:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 유저 탈퇴 실패. 존재하지 않는 사용자입니다.", err)
		case domain.ErrUnableToDeleteUser:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 유저 탈퇴 실패. 유저를 삭제할 수 없습니다.", err)
		default:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 유저 탈퇴 실패. Repository에서 문제 발생", err)
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.DefaultRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		Message:    "✅ 유저 탈퇴 완료",
	})
}
