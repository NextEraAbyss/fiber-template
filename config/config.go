package config

import (
	"os"
)

// Config 保存应用程序配置
type Config struct {
	AppName string
	Port    string
	Env     string
	// 数据库配置
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	// JWT配置
	JWTSecret     string
	JWTExpiration int
}

// Load 返回应用程序配置
func Load() *Config {
	// 默认配置
	config := &Config{
		AppName: getEnv("APP_NAME", "Fiber Template"),
		Port:    getEnv("PORT", "3000"),
		Env:     getEnv("APP_ENV", "development"),
		// 数据库默认值
		DBHost:     getEnv("DB_HOST", "10.0.0.2"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "KW9duzier8sN3AT9"),
		DBName:     getEnv("DB_NAME", "test"),
		// JWT默认值
		JWTSecret:     getEnv("JWT_SECRET", "your_jwt_secret_key"),
		JWTExpiration: 30, // 30分钟
	}

	return config
}

// getEnv 获取环境变量或返回默认值
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
