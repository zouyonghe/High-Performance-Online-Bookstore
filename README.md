<div align="center">

# HPOB · 高性能网上书店

**基于 Go 语言构建的高性能网上书店系统，提供 RESTful API、RBAC 权限控制与开箱即用的 Web 前端。**

[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.25-00ADD8?style=flat-square&logo=go)](https://go.dev/)
[![Go Report Card](https://goreportcard.com/badge/github.com/zouyonghe/High-Performance-Online-Bookstore?style=flat-square)](https://goreportcard.com/report/github.com/zouyonghe/High-Performance-Online-Bookstore)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg?style=flat-square)](LICENSE)
[![API Docs](https://img.shields.io/badge/API-Swagger-85EA2D?style=flat-square&logo=swagger&logoColor=black)](#-api-文档)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](#-贡献指南)

[特性](#-特性) •
[快速开始](#-快速开始) •
[API 文档](#-api-文档) •
[项目结构](#-项目结构) •
[贡献指南](#-贡献指南) •
[许可证](#-许可证)

</div>

---

## 📖 简介

HPOB（High-Performance Online Bookstore）是一个使用 Go 语言编写的在线书店商城服务，基于 [Gin](https://github.com/gin-gonic/gin)、[GORM](https://github.com/go-gorm/gorm)、[Casbin](https://github.com/casbin/casbin)、[Viper](https://github.com/spf13/viper)、[Zap](https://github.com/uber-go/zap) 等成熟生态构建。

系统覆盖用户体系、书籍管理、购物车、订单交易等完整业务流程，内置 Swagger 在线文档与原生 Web 前端，可作为 Go Web 开发的参考工程或直接部署使用。

## ✨ 特性

- **用户体系** — 注册 / 登录 / 注销，JWT 无状态鉴权，bcrypt 密码加密
- **RBAC 权限控制** — 基于 Casbin 的四级角色模型：`admin` / `seller` / `general` / `guest`，策略与代码解耦
- **书籍浏览** — 标题模糊搜索、分类筛选、分页查询，游客亦可访问
- **购物车** — 加购（实时库存校验）、增量移除、一键清空
- **订单交易** — 下单自动清算购物车，支付自动扣减库存（防超卖），支持取消
- **书籍管理** — 卖家与管理员可添加、上下架、删除书籍
- **可观测性** — Zap 结构化日志（lumberjack 滚动切割）、pprof 性能分析、`/state/*` 健康检查（CPU / 内存 / 磁盘）
- **安全基线** — 算法锁定的 JWT 解析、密码哈希不落接口、安全响应头、TLS 支持、依赖漏洞持续扫描
- **开箱即用** — 内置原生 HTML/JS 单页前端与 Swagger UI，启动即可体验

## 🏗 技术栈

| 分层 | 技术选型 |
| --- | --- |
| Web 框架 | [Gin](https://github.com/gin-gonic/gin) |
| ORM / 数据库 | [GORM](https://github.com/go-gorm/gorm) / MySQL · MariaDB |
| 鉴权 / 权限 | [golang-jwt/v5](https://github.com/golang-jwt/jwt) / [Casbin](https://github.com/casbin/casbin) |
| 配置管理 | [Viper](https://github.com/spf13/viper) + [pflag](https://github.com/spf13/pflag) |
| 日志 | [Zap](https://github.com/uber-go/zap) + [lumberjack](https://github.com/natefinch/lumberjack) |
| API 文档 | [Swaggo](https://github.com/swaggo/swag)（Swagger UI） |
| 前端 | 原生 HTML / JavaScript / CSS（零构建依赖） |

系统模块关系图见 [graph.puml](graph.puml) 与 [uml/](uml) 目录。

## 🚀 快速开始

### 环境要求

- Go **1.25+**
- MySQL / MariaDB **5.7+**
- Git

### 1. 克隆仓库

```bash
git clone https://github.com/zouyonghe/High-Performance-Online-Bookstore.git
cd High-Performance-Online-Bookstore
go mod tidy
```

### 2. 初始化数据库

```bash
mysql -uroot -p<PASSWORD> < database/bookstore.sql
```

脚本将创建 `bookstore_server` 数据库、全部数据表，以及内置管理员账户 `admin`。

### 3. 配置

```bash
cp conf/config.example.yaml conf/config.yaml
```

编辑 `conf/config.yaml`，**务必修改数据库密码与 `jwt_secret`**：

```yaml
jwt_secret: <替换为足够长的随机字符串>
db:
  name: bookstore_server
  addr: 127.0.0.1:3306
  username: root
  password: <数据库密码>
```

> ⚠️ `conf/config.yaml`、`conf/server.crt`、`conf/server.key` 已被 `.gitignore` 忽略，**请勿将真实密钥提交到仓库**。完整配置项说明见 [配置参考](#-配置参考)。

### 4. 构建并启动

```bash
go build -o bookstore .
./bookstore            # 或使用脚本 ./server.sh start
```

启动后：

| 入口 | 地址 |
| --- | --- |
| 🖥 Web 前端 | <http://127.0.0.1:8080/> |
| 📑 Swagger 文档 | <http://127.0.0.1:8080/swagger/index.html> |
| 💓 健康检查 | <http://127.0.0.1:8080/state/health> |

### 5. 服务管理

```bash
./server.sh start      # 启动
./server.sh stop       # 停止
./server.sh restart    # 重启
./server.sh version    # 版本信息
```

## 👥 角色与权限

| 能力 | guest | general | seller | admin |
| --- | :-: | :-: | :-: | :-: |
| 浏览 / 搜索书籍 | ✅ | ✅ | ✅ | ✅ |
| 注册 / 登录 | ✅ | ✅ | ✅ | ✅ |
| 购物车 / 下单 / 支付 | | ✅ | | |
| 书籍管理（增删改） | | | ✅ | ✅ |
| 查看已支付订单 | | | ✅ | ✅ |
| 用户管理 | | | | ✅ |

权限策略由 [`conf/policy.csv`](conf/policy.csv) 声明式管理，修改后重启即可生效，无需改动代码。

## 📚 API 文档

启动服务后访问 **Swagger UI**：`http://127.0.0.1:8080/swagger/index.html`

核心接口一览：

| 模块 | 方法 | 路径 | 说明 |
| --- | --- | --- | --- |
| 用户 | `POST` | `/v1/user/register` | 注册普通用户 |
| 用户 | `POST` | `/v1/user/login` | 登录，返回 JWT |
| 用户 | `GET/PUT/DELETE` | `/v1/user/common` | 当前用户信息管理 |
| 用户 | `GET/POST` | `/v1/user/admin` | 用户列表 / 创建卖家（admin） |
| 书籍 | `GET` | `/v1/book` | 书籍列表（`title` / `category` / 分页） |
| 书籍 | `GET` | `/v1/book/:id` | 书籍详情 |
| 书籍 | `POST/PUT/DELETE` | `/v1/book[/:id]` | 书籍管理（seller/admin） |
| 购物车 | `GET/PUT/DELETE` | `/v1/cart` | 查看 / 加购 / 移除 |
| 购物车 | `DELETE` | `/v1/cart/all` | 清空购物车 |
| 订单 | `POST` | `/v1/order` | 从购物车创建订单 |
| 订单 | `PUT` | `/v1/order` | 支付（`accept`）/ 取消（`cancel`） |
| 订单 | `GET` | `/v1/order` | 订单列表（按角色隔离） |
| 状态 | `GET` | `/state/health` 等 | 健康 / 磁盘 / CPU / 内存检查 |

调用示例：

```bash
# 登录获取 token
TOKEN=$(curl -s -X POST http://127.0.0.1:8080/v1/user/login \
  -H 'Content-Type: application/json' \
  -d '{"username":"testuser","password":"test123"}' | jq -r .data.token)

# 搜索书籍（游客亦可）
curl "http://127.0.0.1:8080/v1/book?title=Go&category=编程&pageNum=1&pageSize=10"

# 加入购物车
curl -X PUT http://127.0.0.1:8080/v1/cart \
  -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"BookID":1,"number":2}'

# 下单并支付
ORDER_ID=$(curl -s -X POST http://127.0.0.1:8080/v1/order \
  -H "Authorization: Bearer $TOKEN" | jq -r .data.orderId)
curl -X PUT http://127.0.0.1:8080/v1/order \
  -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' \
  -d "{\"orderId\":$ORDER_ID,\"operation\":\"accept\"}"
```

> 统一响应结构：`{"code": 0, "message": "OK", "data": {...}}`。非零 `code` 为业务错误码，定义见 [`pkg/berror/code.go`](pkg/berror/code.go)。

## ⚙️ 配置参考

`conf/config.yaml` 主要配置项：

| 配置项 | 说明 | 默认值 |
| --- | --- | --- |
| `level` | 运行模式：`debug` / `release` / `test` | `debug` |
| `addr` | HTTP 监听地址 | `:8080` |
| `tls.addr` / `tls.cert` / `tls.key` | HTTPS 监听地址与证书（留空则不启用） | `:8081` |
| `jwt_secret` | JWT 签名密钥（**生产环境必须更换**） | — |
| `max_ping_count` | 启动自检重试次数 | `10` |
| `log.*` | 日志级别、文件输出、滚动策略 | 见示例配置 |
| `db.*` | 主数据库连接 | 见示例配置 |
| `docker_db.*` | 可选的第二数据库（Docker 环境） | 见示例配置 |

## 📁 项目结构

```
.
├── main.go             # 程序入口（pflag / config / database / rbac / router 初始化）
├── server.sh           # 服务管理脚本
├── conf/               # 配置模板、casbin 模型与策略（真实配置不入库）
├── config/             # 配置加载
├── database/           # 数据库连接与初始化 SQL
├── docs/               # Swaggo 生成的 Swagger 文档
├── handler/            # HTTP 处理层：book / cart / order / user / state
├── log/                # 日志初始化与业务日志封装
├── model/              # 领域模型与数据访问：user / book / cart / order
├── permission/         # Casbin RBAC 初始化与鉴权
├── pkg/                # 通用包：auth(加密) / berror(错误码) / token(JWT) / version
├── router/             # 路由注册与中间件（鉴权、安全头、RequestID 等）
├── service/            # 业务逻辑层
├── util/               # 工具函数
├── uml/                # 系统设计 UML 图
└── web/                # 内置单页前端（原生 HTML/JS/CSS）
```

## 🧪 测试

```bash
# 运行全部测试
go test ./...

# 静态检查
go vet ./...

# 依赖漏洞扫描（需安装 govulncheck）
govulncheck ./...
```

测试覆盖 JWT 签发/解析（含算法混淆与过期用例）、bcrypt 加解密、RBAC 策略、统一响应等关键路径。

## 🛣 路线图

- [ ] 订单支付对接真实支付渠道
- [ ] 书籍封面与详情富文本
- [ ] Redis 缓存热点书籍列表
- [ ] Docker Compose 一键部署
- [ ] CI/CD（GitHub Actions）

## 🤝 贡献指南

欢迎提交 Issue 与 Pull Request！

1. Fork 本仓库并创建特性分支：`git checkout -b feat/my-feature`
2. 遵循 [Conventional Commits](https://www.conventionalcommits.org/) 提交规范（如 `feat:`、`fix:`、`docs:`）
3. 提交前确保 `go build ./...`、`go vet ./...`、`go test ./...` 全部通过
4. 改动 API 时请同步执行 `swag init -g main.go --output docs` 更新文档
5. 发起 Pull Request 并描述变更动机与验证方式

安全问题请优先通过邮箱私下报告，而非公开 Issue。

## ❓ 常见问题

**Q：启动时报数据库连接失败？**
确认 MySQL/MariaDB 已启动、`conf/config.yaml` 中账号密码正确，且已执行 `database/bookstore.sql` 完成建表。

**Q：接口返回 403？**
请求路径需要对应角色权限。未携带 token 时按 `guest` 角色处理，请先登录并在请求头携带 `Authorization: Bearer <token>`。

**Q：如何创建 seller / admin 账户？**
注册普通账户后，由 `admin` 通过 `POST /v1/user/admin` 创建卖家账户，或直接在数据库中调整 `tb_users.role`。

## 🙏 致谢

感谢以下优秀的开源项目：

[Go](https://github.com/golang/go) ·
[Gin](https://github.com/gin-gonic/gin) ·
[GORM](https://github.com/go-gorm/gorm) ·
[Viper](https://github.com/spf13/viper) ·
[Zap](https://github.com/uber-go/zap) ·
[Casbin](https://github.com/casbin/casbin) ·
[Swaggo](https://github.com/swaggo/swag)

## 📮 联系

- **Issues**: <https://github.com/zouyonghe/High-Performance-Online-Bookstore/issues>
- **Email**: 1259085392z@gmail.com

## 📄 许可证

本项目基于 [GPL-3.0](LICENSE) 许可证开源。使用者可以自由地使用、修改、再分发本软件源代码，但须遵守许可证相关条款。
