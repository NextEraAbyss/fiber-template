package config

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Config 保存应用程序配置
type Config struct {
	// 应用基础配置
	App struct {
		Name       string
		Env        string
		Debug      bool
		URL        string
		Timezone   string
		Locale     string
		Port       string
		Host       string
		APIPrefix  string
		APITimeout time.Duration
	}

	// 数据库配置
	Database struct {
		Host            string
		Port            string
		User            string
		Password        string
		Name            string
		SSLMode         string
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxLifetime time.Duration
		ConnMaxIdleTime time.Duration
		SlowThreshold   time.Duration
	}

	// Redis配置
	Redis struct {
		Host         string
		Port         string
		Password     string
		DB           int
		PoolSize     int
		MinIdleConns int
		MaxRetries   int
		DialTimeout  time.Duration
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}

	// JWT配置
	JWT struct {
		Secret            string
		Expiration        time.Duration
		RefreshExpiration time.Duration
		Issuer            string
		Audience          string
		Algorithm         string
		HeaderName        string
		HeaderPrefix      string
	}

	// 日志配置
	Log struct {
		Channel    string
		Level      string
		Format     string
		Output     string
		FilePath   string
		MaxSize    int
		MaxBackups int
		MaxAge     int
		Compress   bool
	}

	// 跨域配置
	CORS struct {
		AllowOrigins     []string
		AllowMethods     []string
		AllowHeaders     []string
		ExposeHeaders    []string
		MaxAge           int
		AllowCredentials bool
	}

	// 限流配置
	RateLimit struct {
		Enabled   bool
		Requests  int
		Duration  time.Duration
		Strategy  string
		Whitelist []string
		Blacklist []string
	}

	// 缓存配置
	Cache struct {
		Driver string
		Prefix string
		TTL    time.Duration
		Tags   bool
	}

	// 邮件配置
	Mail struct {
		Mailer      string
		Host        string
		Port        string
		Username    string
		Password    string
		Encryption  string
		FromAddress string
		FromName    string
		LogChannel  string
	}

	// 文件上传配置
	Upload struct {
		Driver         string
		MaxSize        int64
		AllowedTypes   []string
		Path           string
		PublicPath     string
		TempPath       string
		MaxFiles       int
		ImageMaxWidth  int
		ImageMaxHeight int
		ImageQuality   int
	}

	// 安全配置
	Security struct {
		BcryptCost               int
		PasswordMinLength        int
		PasswordMaxLength        int
		PasswordRequireNumbers   bool
		PasswordRequireSymbols   bool
		PasswordRequireUppercase bool
		PasswordRequireLowercase bool
		UsernameMinLength        int
		UsernameMaxLength        int
		UsernameAllowedChars     string
		SessionLifetime          time.Duration
		SessionSecure            bool
		SessionHTTPOnly          bool
		SessionSameSite          string
		CookieLifetime           time.Duration
		CookieSecure             bool
		CookieHTTPOnly           bool
		CookieSameSite           string
	}
}

var config *Config

// Load 返回应用程序配置
func Load() *Config {
	if config != nil {
		return config
	}

	// 加载 .env 文件
	loadEnvFile()

	config = &Config{}

	// 加载应用配置
	loadAppConfig(config)
	// 加载数据库配置
	loadDatabaseConfig(config)
	// 加载Redis配置
	loadRedisConfig(config)
	// 加载JWT配置
	loadJWTConfig(config)
	// 加载日志配置
	loadLogConfig(config)
	// 加载跨域配置
	loadCORSConfig(config)
	// 加载限流配置
	loadRateLimitConfig(config)
	// 加载缓存配置
	loadCacheConfig(config)
	// 加载邮件配置
	loadMailConfig(config)
	// 加载文件上传配置
	loadUploadConfig(config)
	// 加载安全配置
	loadSecurityConfig(config)

	return config
}

// loadEnvFile 加载 .env 文件
func loadEnvFile() {
	// 获取当前工作目录
	workDir, err := os.Getwd()
	if err != nil {
		return
	}

	// 尝试多个可能的 .env 文件位置
	envFiles := []string{
		filepath.Join(workDir, ".env"),
		filepath.Join(workDir, ".env.local"),
		filepath.Join(workDir, ".env.development"),
		filepath.Join(workDir, ".env.production"),
	}

	// 根据环境变量 APP_ENV 确定要加载的文件
	appEnv := os.Getenv("APP_ENV")
	if appEnv != "" {
		envFiles = append([]string{
			filepath.Join(workDir, ".env."+appEnv),
			filepath.Join(workDir, ".env."+appEnv+".local"),
		}, envFiles...)
	}

	// 遍历并加载找到的第一个 .env 文件
	for _, envFile := range envFiles {
		if err := loadEnvFromFile(envFile); err == nil {
			break
		}
	}
}

// loadEnvFromFile 从文件加载环境变量
func loadEnvFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// 跳过空行和注释
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// 解析环境变量
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// 移除引号
		value = strings.Trim(value, `"'`)

		// 如果环境变量未设置，则设置它
		if _, exists := os.LookupEnv(key); !exists {
			os.Setenv(key, value)
		}
	}

	return scanner.Err()
}

// 加载应用配置
func loadAppConfig(c *Config) {
	c.App.Name = getEnv("APP_NAME", "Fiber Template")
	c.App.Env = getEnv("APP_ENV", "development")
	c.App.Debug = getEnvBool("APP_DEBUG", true)
	c.App.URL = getEnv("APP_URL", "http://localhost:3000")
	c.App.Timezone = getEnv("APP_TIMEZONE", "Asia/Shanghai")
	c.App.Locale = getEnv("APP_LOCALE", "zh-CN")
	c.App.Port = getEnv("PORT", "3000")
	c.App.Host = getEnv("HOST", "0.0.0.0")
	c.App.APIPrefix = getEnv("API_PREFIX", "/api/v1")
	c.App.APITimeout = getEnvDuration("API_TIMEOUT", 30*time.Second)
}

// 加载数据库配置
func loadDatabaseConfig(c *Config) {
	c.Database.Host = getEnv("DB_HOST", "localhost")
	c.Database.Port = getEnv("DB_PORT", "5432")
	c.Database.User = getEnv("DB_USER", "postgres")
	c.Database.Password = getEnv("DB_PASSWORD", "postgres")
	c.Database.Name = getEnv("DB_NAME", "fiber_template")
	c.Database.SSLMode = getEnv("DB_SSL_MODE", "disable")
	c.Database.MaxIdleConns = getEnvInt("DB_MAX_IDLE_CONNS", 10)
	c.Database.MaxOpenConns = getEnvInt("DB_MAX_OPEN_CONNS", 100)
	c.Database.ConnMaxLifetime = getEnvDuration("DB_CONN_MAX_LIFETIME", time.Hour)
	c.Database.ConnMaxIdleTime = getEnvDuration("DB_CONN_MAX_IDLE_TIME", 30*time.Minute)
	c.Database.SlowThreshold = getEnvDuration("DB_SLOW_THRESHOLD", 200*time.Millisecond)
}

// 加载Redis配置
func loadRedisConfig(c *Config) {
	c.Redis.Host = getEnv("REDIS_HOST", "localhost")
	c.Redis.Port = getEnv("REDIS_PORT", "6379")
	c.Redis.Password = getEnv("REDIS_PASSWORD", "")
	c.Redis.DB = getEnvInt("REDIS_DB", 0)
	c.Redis.PoolSize = getEnvInt("REDIS_POOL_SIZE", 10)
	c.Redis.MinIdleConns = getEnvInt("REDIS_MIN_IDLE_CONNS", 5)
	c.Redis.MaxRetries = getEnvInt("REDIS_MAX_RETRIES", 3)
	c.Redis.DialTimeout = getEnvDuration("REDIS_DIAL_TIMEOUT", 5*time.Second)
	c.Redis.ReadTimeout = getEnvDuration("REDIS_READ_TIMEOUT", 3*time.Second)
	c.Redis.WriteTimeout = getEnvDuration("REDIS_WRITE_TIMEOUT", 3*time.Second)
}

// 加载JWT配置
func loadJWTConfig(c *Config) {
	c.JWT.Secret = getEnv("JWT_SECRET", "your_jwt_secret_key_here")
	c.JWT.Expiration = getEnvDuration("JWT_EXPIRATION", 24*time.Hour)
	c.JWT.RefreshExpiration = getEnvDuration("JWT_REFRESH_EXPIRATION", 168*time.Hour)
	c.JWT.Issuer = getEnv("JWT_ISSUER", "fiber-template")
	c.JWT.Audience = getEnv("JWT_AUDIENCE", "fiber-template-api")
	c.JWT.Algorithm = getEnv("JWT_ALGORITHM", "HS256")
	c.JWT.HeaderName = getEnv("JWT_HEADER_NAME", "Authorization")
	c.JWT.HeaderPrefix = getEnv("JWT_HEADER_PREFIX", "Bearer")
}

// 加载日志配置
func loadLogConfig(c *Config) {
	c.Log.Channel = getEnv("LOG_CHANNEL", "stack")
	c.Log.Level = getEnv("LOG_LEVEL", "debug")
	c.Log.Format = getEnv("LOG_FORMAT", "json")
	c.Log.Output = getEnv("LOG_OUTPUT", "stdout")
	c.Log.FilePath = getEnv("LOG_FILE_PATH", "./storage/logs")
	c.Log.MaxSize = getEnvInt("LOG_MAX_SIZE", 100)
	c.Log.MaxBackups = getEnvInt("LOG_MAX_BACKUPS", 30)
	c.Log.MaxAge = getEnvInt("LOG_MAX_AGE", 30)
	c.Log.Compress = getEnvBool("LOG_COMPRESS", true)
}

// 加载跨域配置
func loadCORSConfig(c *Config) {
	c.CORS.AllowOrigins = getEnvSlice("CORS_ALLOW_ORIGINS", []string{"http://localhost:3000", "http://localhost:8080"})
	c.CORS.AllowMethods = getEnvSlice("CORS_ALLOW_METHODS", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"})
	c.CORS.AllowHeaders = getEnvSlice("CORS_ALLOW_HEADERS", []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"})
	c.CORS.ExposeHeaders = getEnvSlice("CORS_EXPOSE_HEADERS", []string{"Content-Length", "Content-Range"})
	c.CORS.MaxAge = getEnvInt("CORS_MAX_AGE", 86400)
	c.CORS.AllowCredentials = getEnvBool("CORS_ALLOW_CREDENTIALS", true)
}

// 加载限流配置
func loadRateLimitConfig(c *Config) {
	c.RateLimit.Enabled = getEnvBool("RATE_LIMIT_ENABLED", true)
	c.RateLimit.Requests = getEnvInt("RATE_LIMIT_REQUESTS", 100)
	c.RateLimit.Duration = getEnvDuration("RATE_LIMIT_DURATION", time.Minute)
	c.RateLimit.Strategy = getEnv("RATE_LIMIT_STRATEGY", "ip")
	c.RateLimit.Whitelist = getEnvSlice("RATE_LIMIT_WHITELIST", []string{"127.0.0.1"})
	c.RateLimit.Blacklist = getEnvSlice("RATE_LIMIT_BLACKLIST", []string{})
}

// 加载缓存配置
func loadCacheConfig(c *Config) {
	c.Cache.Driver = getEnv("CACHE_DRIVER", "redis")
	c.Cache.Prefix = getEnv("CACHE_PREFIX", "fiber_template")
	c.Cache.TTL = getEnvDuration("CACHE_TTL", time.Hour)
	c.Cache.Tags = getEnvBool("CACHE_TAGS", true)
}

// 加载邮件配置
func loadMailConfig(c *Config) {
	c.Mail.Mailer = getEnv("MAIL_MAILER", "smtp")
	c.Mail.Host = getEnv("MAIL_HOST", "smtp.gmail.com")
	c.Mail.Port = getEnv("MAIL_PORT", "587")
	c.Mail.Username = getEnv("MAIL_USERNAME", "your_email@gmail.com")
	c.Mail.Password = getEnv("MAIL_PASSWORD", "your_app_password")
	c.Mail.Encryption = getEnv("MAIL_ENCRYPTION", "tls")
	c.Mail.FromAddress = getEnv("MAIL_FROM_ADDRESS", "noreply@yourdomain.com")
	c.Mail.FromName = getEnv("MAIL_FROM_NAME", c.App.Name)
	c.Mail.LogChannel = getEnv("MAIL_LOG_CHANNEL", "stack")
}

// 加载文件上传配置
func loadUploadConfig(c *Config) {
	c.Upload.Driver = getEnv("UPLOAD_DRIVER", "local")
	c.Upload.MaxSize = getEnvInt64("UPLOAD_MAX_SIZE", 10*1024*1024) // 10MB
	c.Upload.AllowedTypes = getEnvSlice("UPLOAD_ALLOWED_TYPES", []string{"image/jpeg", "image/png", "image/gif", "application/pdf"})
	c.Upload.Path = getEnv("UPLOAD_PATH", "./storage/uploads")
	c.Upload.PublicPath = getEnv("UPLOAD_PUBLIC_PATH", "./public/uploads")
	c.Upload.TempPath = getEnv("UPLOAD_TEMP_PATH", "./storage/temp")
	c.Upload.MaxFiles = getEnvInt("UPLOAD_MAX_FILES", 10)
	c.Upload.ImageMaxWidth = getEnvInt("UPLOAD_IMAGE_MAX_WIDTH", 2000)
	c.Upload.ImageMaxHeight = getEnvInt("UPLOAD_IMAGE_MAX_HEIGHT", 2000)
	c.Upload.ImageQuality = getEnvInt("UPLOAD_IMAGE_QUALITY", 85)
}

// 加载安全配置
func loadSecurityConfig(c *Config) {
	c.Security.BcryptCost = getEnvInt("BCRYPT_COST", 10)
	c.Security.PasswordMinLength = getEnvInt("PASSWORD_MIN_LENGTH", 8)
	c.Security.PasswordMaxLength = getEnvInt("PASSWORD_MAX_LENGTH", 100)
	c.Security.PasswordRequireNumbers = getEnvBool("PASSWORD_REQUIRE_NUMBERS", true)
	c.Security.PasswordRequireSymbols = getEnvBool("PASSWORD_REQUIRE_SYMBOLS", true)
	c.Security.PasswordRequireUppercase = getEnvBool("PASSWORD_REQUIRE_UPPERCASE", true)
	c.Security.PasswordRequireLowercase = getEnvBool("PASSWORD_REQUIRE_LOWERCASE", true)
	c.Security.UsernameMinLength = getEnvInt("USERNAME_MIN_LENGTH", 3)
	c.Security.UsernameMaxLength = getEnvInt("USERNAME_MAX_LENGTH", 32)
	c.Security.UsernameAllowedChars = getEnv("USERNAME_ALLOWED_CHARS", "a-z0-9_-.")
	c.Security.SessionLifetime = getEnvDuration("SESSION_LIFETIME", 120*time.Minute)
	c.Security.SessionSecure = getEnvBool("SESSION_SECURE", true)
	c.Security.SessionHTTPOnly = getEnvBool("SESSION_HTTP_ONLY", true)
	c.Security.SessionSameSite = getEnv("SESSION_SAME_SITE", "strict")
	c.Security.CookieLifetime = getEnvDuration("COOKIE_LIFETIME", 120*time.Minute)
	c.Security.CookieSecure = getEnvBool("COOKIE_SECURE", true)
	c.Security.CookieHTTPOnly = getEnvBool("COOKIE_HTTP_ONLY", true)
	c.Security.CookieSameSite = getEnv("COOKIE_SAME_SITE", "strict")
}

// 辅助函数：获取环境变量或返回默认值
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// 辅助函数：获取布尔类型环境变量
func getEnvBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		b, err := strconv.ParseBool(value)
		if err == nil {
			return b
		}
	}
	return defaultValue
}

// 辅助函数：获取整数类型环境变量
func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		i, err := strconv.Atoi(value)
		if err == nil {
			return i
		}
	}
	return defaultValue
}

// 辅助函数：获取int64类型环境变量
func getEnvInt64(key string, defaultValue int64) int64 {
	if value, exists := os.LookupEnv(key); exists {
		i, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return i
		}
	}
	return defaultValue
}

// 辅助函数：获取时间持续时间类型环境变量
func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		d, err := time.ParseDuration(value)
		if err == nil {
			return d
		}
	}
	return defaultValue
}

// 辅助函数：获取字符串切片类型环境变量
func getEnvSlice(key string, defaultValue []string) []string {
	if value, exists := os.LookupEnv(key); exists {
		if value == "" {
			return defaultValue
		}
		return strings.Split(value, ",")
	}
	return defaultValue
}
