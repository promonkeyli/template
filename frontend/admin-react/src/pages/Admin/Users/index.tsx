import { Filter, Search, Users } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { useListUser } from "@/services/api/user/user.ts";
import type {UserListRes} from "@/services/api/model";

export function UsersPage() {
	const { data, isLoading, error } = useListUser({page:1, size: 20});

  if (error) return <div>{error?.message}</div>;
	if (isLoading) return <div>Loading...</div>;


  const d = data?.data?.list as UserListRes[]

	return (
		<div className="space-y-6">
			{/* 搜索和筛选 */}
			<Card>
				<CardHeader>
					<CardTitle>搜索用户</CardTitle>
				</CardHeader>
				<CardContent>
					<div className="flex space-x-4">
						<div className="flex-1">
							<div className="relative">
								<Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-muted-foreground h-4 w-4" />
								<Input placeholder="搜索用户..." className="pl-10" />
							</div>
						</div>
						<Button variant="outline">
							<Filter className="mr-2 h-4 w-4" />
							筛选
						</Button>
					</div>
				</CardContent>
			</Card>

			{/* 用户列表 */}
			<Card>
				<CardHeader>
					<CardTitle className="flex items-center">
						<Users className="mr-2 h-5 w-5" />
						用户列表
					</CardTitle>
				</CardHeader>
				<CardContent>
					<div className="space-y-4">
						{/* 示例用户数据 */}
						{d.map((user) => (
							<div key={user.id} className="flex items-center justify-between p-4 border rounded-lg">
								<div className="flex items-center space-x-4">
									<div className="w-10 h-10 bg-primary rounded-full flex items-center justify-center text-primary-foreground font-medium">
										{user.username?.charAt(0).toUpperCase()}
									</div>
									<div>
										<p className="font-medium">{}</p>
										<p className="text-sm text-muted-foreground">{}</p>
									</div>
								</div>
								<div className="flex items-center space-x-4">
									<span
										className={`px-2 py-1 rounded-full text-xs ${
											user.role === "管理员"
												? "bg-red-100 text-red-800"
												: user.role === "编辑"
													? "bg-blue-100 text-blue-800"
													: "bg-gray-100 text-gray-800"
										}`}
									>
										{user.role}
									</span>
									<span
										className={`px-2 py-1 rounded-full text-xs ${
											user.is_active ? "bg-green-100 text-green-800" : "bg-gray-100 text-gray-800"
										}`}
									>
										{user.is_active ? "活跃" : "非活跃"}
									</span>
									<Button variant="outline" size="sm">
										编辑
									</Button>
								</div>
							</div>
						))}
					</div>
				</CardContent>
			</Card>
		</div>
	);
}
