.PHONY: build run test clean docs swagger fmt deps dev help watch start all

# 默认目标
.DEFAULT_GOAL := help

# 环境变量
APP_NAME = fiber-template
BUILD_DIR = ./build
MAIN_FILE = main.go
PORT = 3000

# 检测操作系统
ifeq ($(OS),Windows_NT)
	RM = if exist $(BUILD_DIR) rd /s /q $(BUILD_DIR)
	MKDIR = if not exist $(BUILD_DIR) mkdir $(BUILD_DIR)
else
	RM = rm -rf $(BUILD_DIR)
	MKDIR = mkdir -p $(BUILD_DIR)
endif

# 构建应用
build:
	@echo "Building $(APP_NAME)..."
	@$(MKDIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo "Build completed"

# 运行应用
run:
	@echo "Running $(APP_NAME)..."
	@go run $(MAIN_FILE)

# 开发模式运行
dev:
	@echo "Running in development mode..."
	@go run $(MAIN_FILE)

# 执行测试
test:
	@echo "Running tests..."
	@go test -v ./...

# 清理构建文件
clean:
	@echo "Cleaning build directory..."
	@$(RM)
	@echo "Cleaned"

# 生成Swagger文档
docs: swagger

# 生成Swagger文档
swagger:
	@echo "Generating Swagger documentation..."
	@swag init
	@swag fmt
	@echo "Swagger documentation generated"

# 格式化代码
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Formatting completed"

# 管理依赖
deps:
	@echo "Tidying dependencies..."
	@go mod tidy
	@echo "Dependencies updated"

# 一键完成所有准备工作并启动服务
start: deps swagger fmt run

# 完整的构建过程：依赖、文档、格式化、构建
all: deps swagger fmt build

# 帮助信息
help:
	@echo "Available commands:"
	@echo "  make build    - Build the application"
	@echo "  make run      - Run the application"
	@echo "  make dev      - Run the application in development mode"
	@echo "  make test     - Run tests"
	@echo "  make clean    - Clean build files"
	@echo "  make docs     - Generate Swagger documentation"
	@echo "  make swagger  - Generate Swagger documentation"
	@echo "  make fmt      - Format code"
	@echo "  make deps     - Update dependencies"
	@echo "  make start    - Update deps, gen swagger, format code and run app"
	@echo "  make all      - Update deps, gen swagger, format code and build"
	@echo "  make help     - Show this help" 