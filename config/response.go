package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Response 统一响应结构
type Response struct {
	Code      int         `json:"code"`      // 业务状态码
	Message   string      `json:"message"`   // 响应消息
	Data      interface{} `json:"data"`      // 响应数据
	Timestamp int64       `json:"timestamp"` // 时间戳
	TraceID   string      `json:"traceId"`   // 追踪ID
}

// Success 成功响应
func Success(c *fiber.Ctx, data interface{}) error {
	return c.JSON(Response{
		Code:      200,
		Message:   "success",
		Data:      data,
		Timestamp: time.Now().Unix(),
		TraceID:   c.GetRespHeader("X-Request-ID"),
	})
}

// Error 错误响应
func Error(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(Response{
		Code:      code,
		Message:   message,
		Data:      nil,
		Timestamp: time.Now().Unix(),
		TraceID:   c.GetRespHeader("X-Request-ID"),
	})
}

// ValidationError 验证错误响应
func ValidationErrors(c *fiber.Ctx, message string) error {
	return Error(c, 400, message)
}

// UnauthorizedError 未授权错误响应
func UnauthorizedError(c *fiber.Ctx) error {
	return Error(c, 401, "未授权访问")
}

// ForbiddenError 禁止访问错误响应
func ForbiddenError(c *fiber.Ctx) error {
	return Error(c, 403, "禁止访问")
}

// NotFoundError 资源不存在错误响应
func NotFoundError(c *fiber.Ctx) error {
	return Error(c, 404, "资源不存在")
}

// ServerError 服务器错误响应
func ServerError(c *fiber.Ctx) error {
	return Error(c, 500, "服务器内部错误")
}

// BadRequestError 请求错误响应
func BadRequestError(c *fiber.Ctx, message string) error {
	return Error(c, 400, message)
}

// ConflictError 冲突错误响应
func ConflictError(c *fiber.Ctx, message string) error {
	return Error(c, 409, message)
}

// TooManyRequestsError 请求过多错误响应
func TooManyRequestsError(c *fiber.Ctx) error {
	return Error(c, 429, "请求过于频繁")
}

// ServiceUnavailableError 服务不可用错误响应
func ServiceUnavailableError(c *fiber.Ctx) error {
	return Error(c, 503, "服务暂时不可用")
}

// GatewayTimeoutError 网关超时错误响应
func GatewayTimeoutError(c *fiber.Ctx) error {
	return Error(c, 504, "网关超时")
}
