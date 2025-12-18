package main

import (
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	// 1. 执行 swag fmt 命令
	if err := runCommand("swag", "fmt"); err != nil {
		slog.Error(err.Error())
		return
	}

	// 2.openapi注释里先添加程序入口目录
	dirs := []string{"./cmd/server"}

	// 3. 添加internal/app 下的所有子目录
	appModules, err := scanAppModules("./internal/app")
	if err != nil {
		slog.Error("扫描模块目录失败: " + err.Error())
		return
	}
	dirs = append(dirs, appModules...)

	// 4. 添加其他需要扫描的目录：例如此处有定义公共的响应体，以及响应code
	dirs = append(dirs, "./internal/pkg/http")

	dir := strings.Join(dirs, ",")

	// 5. 执行 swag init 命令, 输出 openapi 文档到指定目录
	if err := runCommand("swag", "init", "--dir", dir, "-o", "./api/openapi"); err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info("openAPI 文档生成成功")
}

// scanAppModules 扫描指定目录下的所有子目录
func scanAppModules(basePath string) ([]string, error) {
	var modules []string

	entries, err := os.ReadDir(basePath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			// 构建相对路径
			modulePath := filepath.Join(basePath, entry.Name())
			modules = append(modules, modulePath)
		}
	}

	return modules, nil
}

func runCommand(name string, args ...string) error {
	// 创建命令对象
	cmd := exec.Command(name, args...)
	// 获取命令的标准输出和标准错误
	output, err := cmd.CombinedOutput()
	slog.Debug(string(output))
	return err
}
