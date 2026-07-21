# HPOB 高性能网上书店

A high performance online bookstore system.

## Introduction 介绍

一个基于Gin、gorm、viper、zap等库的web服务器，实现了网上书店相关接口。

### Summary 概要

使用go语言编写的，基于gin、gorm、viper、zap等模块的在线书店商城服务。

### Features 特性

路由高效，日志高性能，配置读取便捷，代码通俗易懂。

- 用户体系：注册、登录、JWT 鉴权，admin / seller / general / guest 四种角色（casbin RBAC）
- 书籍浏览：标题模糊搜索、分类筛选、分页
- 购物车：加购（含库存校验）、移除、清空
- 订单：下单、支付（自动扣减库存）、取消
- 书籍管理：卖家/管理员添加、上下架、删除书籍
- 自带简单前端（web/ 目录，原生 HTML/JS 单页应用）与 Swagger 在线文档

## Requirements 必要条件

系统支持Go语言环境，已经安装mariadb或mysql数据库，其它模块依赖参考go.mod文件。

## Configuration 配置

复制示例配置文件并修改其中的配置（真实配置已被 git 忽略，请勿提交真实密钥）：

```bash
cp conf/config.example.yaml conf/config.yaml
```

在conf/config.yaml文件中修改配置。

```yaml
level: debug                 # 开发模式, debug, release, test
name: bookstore-server       # API服务的名字
addr: :8080                  # HTTP绑定端口
url: http://127.0.0.1:8080   # pingServer函数请求的API服务器的ip:port
max_ping_count: 10           # pingServer函数try的次数
jwt_secret: Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5 #token密钥
tls:
  addr: :8081                # https地址
  cert: conf/server.crt      # 证书地址
  key: conf/server.key       # 私钥文件
log:
  log_level: DEBUG           # 日志级别，可选DEBUG, INFO, WARN, ERROR, DPANIC, PANIC, FATAL
  file_output: true          # 是否在文件中输出日志
  log_file: blog/server.log  # 输出日志文件路径
  max_backups: 5             # 日志文件最大保留数量
  max_size: 1                # 日志文件最大大小，单位 MB
  max_age: 30                # 日志文件最大存储时间，单位 天
  compress: false            # 是否启用日志压缩
db:
  name: bookstore_server     # 数据库名称
  addr: 127.0.0.1:3306       # 数据库所在IP及端口号
  username: root             # 数据库登陆用户
  password: aa11bb22cc33     # 数据库登陆密码
docker_db:
  name: bookstore_server
  addr: 127.0.0.1:3306
  username: root
  password: aa11bb22cc33
```

## Installation 安装

```bash
#克隆存储库
git clone https://github.com/zouyonghe/High-Performance-Online-Bookstore.git

#进入项目目录
cd High-Performance-Online-Bookstore

#安装依赖库
go mod tidy
```

## Usage 用法

```bash
#初始化数据库

cd database
mysql -uroot -pPASSWORD -e "
source bookstore.sql
quit
"
cd ..

#编译程序

go build

#启动服务

./server.sh start

#访问前端页面

#浏览器打开 http://127.0.0.1:8080/ （自动跳转到 /web/index.html）

#访问API文档

#浏览器打开 http://127.0.0.1:8080/swagger/index.html

#终止服务

./server.sh stop

#重启服务

./server.sh restart

#版本信息

./server.sh version
```

### API 说明

列表类接口通过 query string 传参，例如：

```bash
#书籍列表（标题模糊搜索、分类筛选、分页）
curl "http://127.0.0.1:8080/v1/book?title=Go&category=编程&pageNum=1&pageSize=10"

#订单列表（需携带 token）
curl "http://127.0.0.1:8080/v1/order?pageNum=1&pageSize=10" \
  -H "Authorization: Bearer <token>"
```

其余接口参数参考 Swagger 文档。

## Struct 项目结构

#### main

程序入口。

#### conf

配置文件，yaml文档及TLS证书和私钥，基于角色的权限控制配置等。

#### config

服务配置获取。

#### database

数据库连接配置，初始化脚本。

#### log

日志配置获取。

#### router

路由与中间件。

> - router.go 路由函数及中间件注册
> - middleware 中间件

#### handler

路由处理函数，包括用户路由处理函数等。

> - handler.go 响应结构体，响应函数
> - state 状态查询响应
> - user 用户请求响应

#### model

服务主体对象类型及方法。

> - model.go 其他数据模型
> - user.go 用户模型和方法
> - book.go 图书模型和方法
> - cart.go 购物车模型和方法
> - order.go 订单模型和方法

#### service

复杂业务逻辑函数。

#### util

工具包。

#### pkg

其他功能包，认证、错误码、状态、token、版本等。

## Development 开发文档

待补充。

## FAQ - 常见问题

## Support 支持

如果需要反馈问题，可以直接提交 [Issues](https://github.com/zouyonghe/High-Performance-Online-Bookstore/issues) 或联系邮箱。

### Dos 文档

### Contact 联系

name: buding

email: 1259085392z@gamil.com

## Authors and acknowledgment 贡献者和感谢

The [Go](https://github.com/golang/go) programming language

[Gin](https://github.com/gin-gonic/gin) is a HTTP web framework written in Golang. 

[Gorm](https://github.com/go-gorm/gorm) is a fantastic ORM library for Golang, aims to be developer friendly.

[Viper](https://github.com/spf13/viper) is a complete configuration solution for Go applications including 12-Factor apps.

[Casbin](https://github.com/casbin/casbin) is a powerful and efficient open-source access control library for Golang projects.

## License 版权信息

本软件基于 [GPLv3](https://github.com/zouyonghe/High-Performance-Online-Bookstore/blob/main/LICENSE) 协议开源，使用者可以自由的使用、修改、再分发此软件源代码，但须遵守许可证相关条例。
