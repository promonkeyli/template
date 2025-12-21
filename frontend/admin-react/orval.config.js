import { defineConfig } from "orval";

export default defineConfig({
	"mall-service": {
		input: "http://127.0.0.1:8081/swagger/doc.json",
		output: {
			target: "./src/services/api",
			client: "react-query", // 生成 React Query Hooks
			schemas: "./src/services/api/model",
			mode: "tags-split", // 按 tag 分割文件
			mock: true, // msw mock 功能开启
			override: {
				mutator: {
					// 自定义请求适配器
					path: "./src/services/request.ts",
					name: "customInstance",
				},
			},
		},
	},
});
