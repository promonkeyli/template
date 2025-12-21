import path from "node:path";
import { fileURLToPath } from "node:url";
import tailwindcss from "@tailwindcss/vite";
import { tanstackRouter } from "@tanstack/router-plugin/vite";
import react from "@vitejs/plugin-react";
import { defineConfig, loadEnv } from "vite";

// // ESM中 获取 __dirname
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

export default defineConfig(({ mode }) => {
	// 1. 设置 env 目录路径
	const envDir = path.resolve(__dirname, "env");
	// 2. 加载环境变量 (参数: mode, 目录, 前缀)
	const env = loadEnv(mode, envDir, "");

	return {
		// 你的原有配置
		plugins: [
			tanstackRouter({
				target: "react",
				autoCodeSplitting: true,
				generatedRouteTree: "./src/router.ts", // 手动指定routeTree.gen.ts 生成的位置以及名称
			}),
			react({
				babel: {
					plugins: ["babel-plugin-react-compiler"],
				},
			}),
			tailwindcss(),
		],

		// 保持你的 envDir 配置
		envDir: envDir,

		resolve: {
			alias: {
				"@": path.resolve(__dirname, "./src"),
			},
		},

		// --- 👇 重点：这里可以解决你之前 Orval/Axios 的报错 ---
		// 通过 define 将读取到的环境变量注入到全局，模拟 process.env
		define: {
			"process.env": JSON.stringify(env),
		},
	};
});
