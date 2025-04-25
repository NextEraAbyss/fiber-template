package controller

import (
	"strconv"

	"github.com/NextEraAbyss/fiber-template/app/service"
	"github.com/gofiber/fiber/v2"
)

// UserController 用户控制器
type UserController struct {
	userService *service.UserService
}

// NewUserController 创建新的用户控制器实例
func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// GetUsers 获取用户列表
// @Summary 获取用户列表
// @Description 获取用户列表，支持分页和筛选
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Param username query string false "用户名"
// @Param email query string false "邮箱"
// @Param role query string false "角色"
// @Param is_active query string false "是否激活"
// @Success 200 {object} fiber.Map
// @Router /api/v1/users [get]
func (c *UserController) GetUsers(ctx *fiber.Ctx) error {
	// 获取查询参数
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size", "10"))

	params := service.UserQueryParams{
		Page:     page,
		PageSize: pageSize,
		Username: ctx.Query("username"),
		Email:    ctx.Query("email"),
		Role:     ctx.Query("role"),
		IsActive: ctx.Query("is_active"),
	}

	// 获取用户列表
	users, pagination, err := c.userService.GetUsers(params)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"users":      users,
		"pagination": pagination,
	})
}

// GetUser 获取单个用户
// @Summary 获取单个用户
// @Description 通过ID获取单个用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} model.User
// @Router /api/v1/users/{id} [get]
func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "无效的用户ID")
	}

	user, err := c.userService.GetUserByID(uint(id))
	if err != nil {
		if err == service.ErrUserNotFound {
			return fiber.NewError(fiber.StatusNotFound, "用户不存在")
		}
		return err
	}

	return ctx.JSON(user)
}
