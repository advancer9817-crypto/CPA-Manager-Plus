# CPA Manager Plus 部署迭代记录

## 2026-05-27: 重新构建 & 部署

### 变更摘要
- 前端 `apps/web` 重新构建（vite singlefile 2.9MB）
- CPAM binary `apps/manager-server` 重新编译（stripped, 静态链接）
- 同步前端到 `.deploy/management.html` 和 `CLIProxyAPI/static/management.html`

### 影响范围
| 文件 | 说明 |
|------|------|
| `apps/web/dist/index.html` | 新构建的前端 |
| `.deploy/management.html` | 同步为 CPAM 运行时面板 |
| `.deploy/cpa-manager-plus` | 重新编译的 manager-server 二进制 |

### 配套部署
- systemd user services: `cpa-proxy`, `cpa-manager-plus`
- `~/.local/bin/cpa` 一键管理命令

### 回滚
1. 恢复 `.deploy/cpa-manager-plus` 到旧版本
2. 恢复 `.deploy/management.html` 到旧版本
3. 重新加载 systemd: `systemctl --user daemon-reload && cpa restart`

## [迭代] 2026-06-23 — 升级到 v1.8.1

- 从 upstream 同步并合并了 `v1.8.1` 标签的内容
- 重新构建了前端 (`pnpm build`)
- 部署了最新构建的 `apps/web/dist/index.html` 到 `.deploy/management.html`
- 重启了 `cpa-manager-plus` 用户服务

## [迭代] 2026-07-09 — 同步 upstream 至 v1.10.5 + 版本检查后端代理

### 变更摘要
合并 upstream v1.8.1 → v1.10.5（41 个提交），同时新增本地定制功能。

#### upstream 合入
- v1.10.4 / v1.10.5 发布版本
- 配额插件更新：xAI 周配额摘要、release version picker
- Provider 优先级内联编辑
- 监控 filter labels / quota tooltips 修复
- Provider recent status overlap 修复

#### 本地定制
- **版本检查走后端代理**：前端不再直连 GitHub API 查询 CPA-Manager-Plus 最新版本，改为调用后端 `/v0/management/latest-manager-version` 端点，避免浏览器匿名 GitHub API 限流 403
  - 新增 `apps/manager-server/internal/http/controller/system/latest_manager_version.go`
  - 路由注册 `router.go` 新增 `/v0/management/latest-manager-version`
  - 前端 `version.ts` 移除 axios 直连 GitHub，改用 `apiClient`
- **构建日期占位符处理**：`VersionCard.tsx` 新增 `BUILD_DATE_PLACEHOLDERS` 集合，对 `unknown`/`dev`/`none` 等占位符显示为"未知"而非 `Invalid Date`
- **管理面板更新**：`management.html` 内嵌最新构建的控制面板
- 新增 `encrypt_key` / `decrypt_key` 工具命令

### 提交数
41 个未推送提交 + 本地未提交改动
