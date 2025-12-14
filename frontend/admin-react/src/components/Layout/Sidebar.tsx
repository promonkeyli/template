import { useState } from 'react'
import { useNavigate, useLocation } from '@tanstack/react-router'
import {
  LayoutDashboard,
  Users,
  Settings,
  BarChart3,
  ChevronLeft,
  ChevronRight,
  Home,
  User,
  Bell,
  Mail
} from 'lucide-react'
import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'
import { cn } from '@/lib/utils'

interface MenuItem {
  id: string
  label: string
  icon: React.ComponentType<{ className?: string }>
  path: string
  children?: MenuItem[]
}

const menuItems: MenuItem[] = [
  {
    id: 'dashboard',
    label: '仪表盘',
    icon: LayoutDashboard,
    path: '/admin'
  },
  {
    id: 'users',
    label: '用户管理',
    icon: Users,
    path: '/admin/users',
    children: [
      { id: 'user-list', label: '用户列表', icon: User, path: '/admin/users' },
      { id: 'user-roles', label: '角色管理', icon: Settings, path: '/admin/users/roles' }
    ]
  },
  {
    id: 'analytics',
    label: '数据分析',
    icon: BarChart3,
    path: '/admin/analytics'
  },
  {
    id: 'system',
    label: '系统设置',
    icon: Settings,
    path: '/admin/system',
    children: [
      { id: 'system-config', label: '系统配置', icon: Settings, path: '/admin/system/config' },
      { id: 'notifications', label: '通知管理', icon: Bell, path: '/admin/system/notifications' },
      { id: 'messages', label: '消息管理', icon: Mail, path: '/admin/system/messages' }
    ]
  }
]

interface SidebarProps {
  collapsed: boolean
  onToggle: () => void
}

export function Sidebar({ collapsed, onToggle }: SidebarProps) {
  const navigate = useNavigate()
  const location = useLocation()
  const [expandedItems, setExpandedItems] = useState<string[]>(['users', 'content', 'system'])

  const handleMenuClick = (item: MenuItem) => {
    if (item.children) {
      // 如果有子菜单，切换展开状态
      setExpandedItems(prev =>
        prev.includes(item.id)
          ? prev.filter(id => id !== item.id)
          : [...prev, item.id]
      )
    } else {
      // 如果没有子菜单，直接导航
      navigate({ to: item.path })
    }
  }

  const isActive = (path: string) => {
    return location.pathname === path
  }

  const isExpanded = (id: string) => {
    return expandedItems.includes(id)
  }

  return (
    <div className={cn(
      "flex flex-col bg-card border-r border-border transition-all duration-300",
      collapsed ? "w-16" : "w-64"
    )}>
      {/* 侧边栏头部 */}
      <div className="flex items-center justify-between p-4 border-b border-border">
        {!collapsed && (
          <div className="flex items-center space-x-2">
            <div className="w-8 h-8 bg-primary rounded-lg flex items-center justify-center">
              <Home className="w-5 h-5 text-primary-foreground" />
            </div>
            <span className="font-semibold text-lg">管理后台</span>
          </div>
        )}
        <Button
          variant="ghost"
          size="sm"
          onClick={onToggle}
          className="h-8 w-8 p-0"
        >
          {collapsed ? (
            <ChevronRight className="h-4 w-4" />
          ) : (
            <ChevronLeft className="h-4 w-4" />
          )}
        </Button>
      </div>

      {/* 菜单项 */}
      <nav className="flex-1 p-4 space-y-2">
        {menuItems.map((item) => (
          <div key={item.id}>
            <Button
              variant={isActive(item.path) ? "secondary" : "ghost"}
              className={cn(
                "w-full justify-start h-10",
                collapsed ? "px-2" : "px-3",
                isActive(item.path) && "bg-secondary text-secondary-foreground"
              )}
              onClick={() => handleMenuClick(item)}
            >
              <item.icon className={cn("h-4 w-4", !collapsed && "mr-3")} />
              {!collapsed && (
                <span className="flex-1 text-left">{item.label}</span>
              )}
              {!collapsed && item.children && (
                <ChevronRight
                  className={cn(
                    "h-4 w-4 transition-transform",
                    isExpanded(item.id) && "rotate-90"
                  )}
                />
              )}
            </Button>

            {/* 子菜单 */}
            {!collapsed && item.children && isExpanded(item.id) && (
              <div className="ml-6 mt-2 space-y-1">
                {item.children.map((child) => (
                  <Button
                    key={child.id}
                    variant={isActive(child.path) ? "secondary" : "ghost"}
                    size="sm"
                    className={cn(
                      "w-full justify-start h-8 text-sm",
                      isActive(child.path) && "bg-secondary text-secondary-foreground"
                    )}
                    onClick={() => navigate({ to: child.path })}
                  >
                    <child.icon className="h-3 w-3 mr-2" />
                    {child.label}
                  </Button>
                ))}
              </div>
            )}
          </div>
        ))}
      </nav>

      {/* 侧边栏底部 */}
      <div className="p-4 border-t border-border">
        <Separator className="mb-4" />
        {!collapsed && (
          <div className="text-xs text-muted-foreground text-center">
            v1.0.0
          </div>
        )}
      </div>
    </div>
  )
}
