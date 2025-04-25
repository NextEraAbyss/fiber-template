package config

import (
	"regexp"
	"strings"
	"unicode"
)

// Sanitize 清理输入字符串，移除危险字符
func Sanitize(input string) string {
	// 移除所有HTML标签
	re := regexp.MustCompile("<[^>]*>")
	sanitized := re.ReplaceAllString(input, "")

	// 移除所有可能在脚本中使用的特殊字符
	sanitized = strings.Map(func(r rune) rune {
		switch r {
		case '<', '>', '&', '"', '\'', '`', ';', '(', ')', '{', '}':
			return -1
		default:
			return r
		}
	}, sanitized)

	return strings.TrimSpace(sanitized)
}

// SanitizeUsername 清理用户名，只允许字母、数字、下划线和中划线
func SanitizeUsername(username string) string {
	// 只保留字母、数字、下划线和中划线
	sanitized := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-' {
			return r
		}
		return -1
	}, username)

	return strings.TrimSpace(sanitized)
}

// SanitizeEmail 清理并验证电子邮件格式
func SanitizeEmail(email string) (string, bool) {
	// 移除所有空白字符
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	// 验证电子邮件格式
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	isValid := re.MatchString(email)

	return email, isValid
}

// EscapeHTML 转义HTML特殊字符
func EscapeHTML(input string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		"\"", "&quot;",
		"'", "&#39;",
	)
	return replacer.Replace(input)
}
