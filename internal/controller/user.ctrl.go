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

func (c *UserController) CreateUserProfile(ctx *fiber.Ctx) error {
	createUserProfilePayload := new(domain.UserCreateUserProfileReq)
	createUserProfilePayload.UserUUID = middleware.GetUUIDFromMiddleware(ctx)

	if err := c.userService.CreateUserProfile(createUserProfilePayload); err != nil {
		log.Printf("%v", err)
		switch err {
		case model.ErrNotFound:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 유저 프로필 생성 실패. 존재하지 않는 사용자 프로필을 생성 시도 중입니다.", err)
		default:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 유저 프로필 생성 실패. Repository에서 문제 발생", err)
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.DefaultRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		Message:    "✅ 유저 프로필 생성 완료",
	})
}

func (c *UserController) GetUserProfile(ctx *fiber.Ctx) error {
	profileTag := ctx.Params("userTag")
	if len(profileTag) <= 0 {
		return utils.CreateErrorRes(ctx, fiber.StatusBadRequest, "❌ 유저 조회 실패. Binding 과정에서 문제 발생", domain.ErrMissingParams)
	}

	userProfile, err := c.userService.GetUserUUIDByTag(profileTag)
	if err != nil {
		log.Printf("%v", err)
		switch err {
		case model.ErrNotFound:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 유저 조회 실패. 존재하지 않는 사용자입니다.", err)
		default:
			return utils.CreateErrorRes(ctx, fiber.StatusInternalServerError, "❌ 유저 조회 실패. Repository에서 문제 발생", err)
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(domain.UserGetUserProfileRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		Message:    "✅ 유저 프로필 조회 완료",
		Profile:    userProfile,
	})
}

func (c *UserController) UpdateUserProfile(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(domain.DefaultRes{
		IsError:    false,
		StatusCode: fiber.StatusOK,
		Message:    "✅ 유저 프로필 수정 완료",
	})
}
