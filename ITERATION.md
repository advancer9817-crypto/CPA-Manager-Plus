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
