package config

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidationError 表示验证错误
type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

// Validator 是验证器的实例
var Validator = validator.New(validator.WithRequiredStructEnabled())

// ValidateStruct 验证结构体
func ValidateStruct(s interface{}) []*ValidationError {
	// 初始化验证器
	validate := Validator

	// 注册自定义验证函数
	registerCustomValidations(validate)

	// 执行验证
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	// 处理验证错误
	errors := err.(validator.ValidationErrors)
	var validationErrors []*ValidationError

	for _, err := range errors {
		// 提取错误信息
		validationError := &ValidationError{
			Field: toSnakeCase(err.Field()),
			Tag:   err.Tag(),
			Value: err.Param(),
		}
		validationErrors = append(validationErrors, validationError)
	}

	return validationErrors
}

// 注册自定义验证函数
func registerCustomValidations(validate *validator.Validate) {
	// 安全的用户名验证: 只允许字母、数字和下划线，长度3-32
	validate.RegisterValidation("safename", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		if val == "" {
			return true // 如果字段是可选的
		}

		match, _ := regexp.MatchString(`^[a-zA-Z0-9_]{3,32}$`, val)
		return match
	})

	// 安全文本验证: 过滤特殊字符或潜在的危险符号
	validate.RegisterValidation("safetext", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		if val == "" {
			return true // 如果字段是可选的
		}

		// 检查是否包含潜在危险符号
		dangerousPatterns := []string{"<script", "javascript:", "onerror=", "onload=", "--", "DROP", "DELETE FROM"}
		for _, pattern := range dangerousPatterns {
			if strings.Contains(strings.ToLower(val), strings.ToLower(pattern)) {
				return false
			}
		}
		return true
	})
}

// 将驼峰命名转换为蛇形命名
func toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// ValidateJSON 验证JSON输入
func ValidateJSON(s interface{}) bool {
	return reflect.ValueOf(s).Kind() == reflect.Struct ||
		reflect.ValueOf(s).Kind() == reflect.Map
}

// SanitizeString 清理字符串，移除潜在危险字符
func SanitizeString(input string) string {
	// 移除HTML标签
	re := regexp.MustCompile("<[^>]*>")
	sanitized := re.ReplaceAllString(input, "")

	// 移除其他危险字符
	dangerousChars := []string{"--", ";", "/*", "*/", "@@", "@", "char(", "exec(", "union", "select", "insert", "drop", "update", "delete"}
	for _, char := range dangerousChars {
		sanitized = strings.ReplaceAll(strings.ToLower(sanitized), char, "")
	}

	return sanitized
}
