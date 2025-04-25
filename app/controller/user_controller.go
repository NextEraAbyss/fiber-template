package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/fiber-template/app/model"
	"github.com/yourusername/fiber-template/app/service"
)

// GetUsers 返回所有用户
//	@Summary		获取所有用户
//	@Description	获取系统中所有用户的列表
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"包含用户列表的响应"
//	@Failure		500	{object}	map[string]interface{}	"服务器错误"
//	@Router			/users [get]
func GetUsers(c *fiber.Ctx) error {
	users, err := service.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    users,
	})
}

// GetUser 通过ID返回特定用户
//	@Summary		获取单个用户
//	@Description	通过ID获取特定用户的详细信息
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string					true	"用户ID"
//	@Success		200	{object}	map[string]interface{}	"包含用户信息的响应"
//	@Failure		404	{object}	map[string]interface{}	"用户未找到"
//	@Router			/users/{id} [get]
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := service.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "用户未找到",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

// CreateUser 创建新用户
//	@Summary		创建用户
//	@Description	创建一个新用户
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.User				true	"用户信息"
//	@Success		201		{object}	map[string]interface{}	"创建成功的响应"
//	@Failure		400		{object}	map[string]interface{}	"请求体错误"
//	@Failure		500		{object}	map[string]interface{}	"服务器错误"
//	@Router			/users [post]
func CreateUser(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "无效的请求体",
		})
	}

	newUser, err := service.CreateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    newUser,
	})
}

// UpdateUser 通过ID更新用户
//	@Summary		更新用户
//	@Description	通过ID更新特定用户的信息
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string					true	"用户ID"
//	@Param			user	body		model.User				true	"用户信息"
//	@Success		200		{object}	map[string]interface{}	"更新成功的响应"
//	@Failure		400		{object}	map[string]interface{}	"请求体错误"
//	@Failure		500		{object}	map[string]interface{}	"服务器错误"
//	@Router			/users/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "无效的请求体",
		})
	}

	updatedUser, err := service.UpdateUser(id, &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    updatedUser,
	})
}

// DeleteUser 通过ID删除用户
//	@Summary		删除用户
//	@Description	通过ID删除特定用户
//	@Tags			用户管理
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string					true	"用户ID"
//	@Success		200	{object}	map[string]interface{}	"删除成功的响应"
//	@Failure		500	{object}	map[string]interface{}	"服务器错误"
//	@Router			/users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := service.DeleteUser(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "用户已删除",
	})
}
