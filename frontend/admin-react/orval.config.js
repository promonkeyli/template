import { defineConfig } from "orval";

// 配置：https://orval.dev/reference/configuration/overview
export default defineConfig({
	"mall-service": {
		input: "http://127.0.0.1:8081/swagger/doc.json",
		output: {
			target: "./src/services/api",
			client: "react-query", // Tankstack Query 作为客户端
			schemas: "./src/services/api/model", // 模型文件生成路径声明
			mode: "tags-split", // 按 tag 分割文件
			prettier: false, // 关闭自带的prettier，使用biome进行文档格式化
			biome: true, // 开启 biome
			mock: false, // msw mock 功能暂时关闭
			override: {
				mutator: {
					// 自定义请求适配器
					path: "./src/services/core/request.ts",
					name: "customInstance",
				},
			},
		},
		// hooks: {
		// 	afterAllFilesWrite: "npx @biomejs/biome check --write", // biome: true 配置开启时，不需要配置该项
		// },
	},
});
