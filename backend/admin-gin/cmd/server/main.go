//	@title			Mall API
//	@version		1.0
//	@description	Mall API 服务接口文档
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8081
//	@BasePath	/

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.

package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"mall-api/configs"
	"mall-api/internal/boot"
)

func main() {

	// 1.从环境变量获取运行模式
	runMode := os.Getenv("APP_ENV_MODE")
	env := "dev"
	if runMode != "" {
		env = runMode
	}

	// 2. 依据环境变量初始化系统配置
	cfg, err := configs.InitConfig(env)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// 3. 构造应用
	app, appErr := boot.NewApp(cfg)
	if appErr != nil {
		slog.Error(appErr.Error())
		os.Exit(1)
	}

	// 4.注入依赖
	boot.Register(app.Ge, app.Db, app.Rdb, app.Jt)

	// 5. 端口打印
	fmt.Printf("【%s】service is running on port: %d \n\n", strings.ToUpper(cfg.App.Name), cfg.Server.Port)

	// 6. 运行 http 服务
	app.Se.ListenAndServe()
}
