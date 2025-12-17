# 项目名称
APP_NAME=mall-api

# 主程序入口文件
MAIN=./cmd/server/main.go

# OpenAPI 文档生成工具入口文件
OPENAPI=./cmd/openapi/main.go

# 数据库迁移
MIGRATE= ./cmd/migrate/main.go

# Go 工具安装目录（用于找到 swag 等工具）
GOBIN=$(shell go env GOPATH)/bin

# 运行服务器
run:
	go run $(MAIN)

# 生成 Open API 文档
# 注意：需要临时添加 GOBIN 到 PATH，以便找到 swag 命令
swag:
	@echo "生成 API 文档..."
	@PATH=$(GOBIN):$$PATH go run $(OPENAPI)

# web 服务构建
build:
	@echo "正在构建..."
	go build -o bin/$(BINARY_NAME) cmd/server/main.go

# postgre 数据库迁移
migrate:
	@echo "postgre数据库迁移..."
	go run $(MIGRATE)
