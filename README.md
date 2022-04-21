# Jinshuzhai-Bookstore 锦书斋网上书店

An online bookstore system.

## Introduction 介绍

一个基于Gin、gorm、viper、zap等库的web服务器，实现了网上书店相关接口。

### Summary 概要

使用go语言编写的，基于gin、gorm、viper、zap等模块的在线书店商城后端服务。

### Features 特性

路由高效性，日志高性能，配置读取高便捷性，代码通俗易懂。

## Requirements 必要条件

系统支持Go语言环境，已经安装mariadb或mysql数据库，其它模块依赖参考go.mod文件。

## Configuration 配置

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
git clone https://github.com/zouyonghe/Jinshuzhai-Bookstore.git

#进入项目目录
cd Jinshuzhai-Bookstore

#安装依赖库
go mod tidy
```

## Usage 用法

```bash
#初始化数据库
cd sql
mysql -uroot -pPASSWORD -e "
source db.sql
quit
"
cd ..

#启动服务

./server.sh start

#终止服务

./server.sh stop



```

## Struct 项目结构

#### main

程序入口。

#### config

服务配置获取。

#### log

日志配置获取。

#### router

路由与中间件。

#### handler

路由处理函数，包括用户路由处理函数等。

#### model

对象关系映射，建立数据库连接，CRUD等。

#### sql

数据库初始化。

#### util

工具包，支持生成短ID、获取请求ID等。

#### pkg

其他功能包，支持认证、错误码、状态、token、版本等。

## Development 开发文档

## Changelog 更新日志

| 日期        | 事项                 | 开发者    |
|:---------:|:------------------:|:------:|
| 2022-4-14 | 初始化项目，完成配置和日志记录代码  | buding |
| 2022-4-15 | 完成版本信息，调整配置和日志记录代码 | buding |
| 2022-4-21 | 完成用户接口设计，接口文档生成    | buding |

## FAQ - 常见问题

## Support 支持

如果需要反馈问题，可以直接提交 [Issues](https://github.com/zouyonghe/Jinshuzhai-Bookstore/issues) 或联系邮箱。

### Dos 文档

### Contact 联系

name: buding

email: 1259085392z@gamil.com

## Authors and acknowledgment 贡献者和感谢

The [Go](https://github.com/golang/go) programming language

[Gin](https://github.com/gin-gonic/gin) is a HTTP web framework written in Golang. 

[Gorm](https://github.com/go-gorm/gorm) is a fantastic ORM library for Golang, aims to be developer friendly.

[Viper](https://github.com/spf13/viper) is a complete configuration solution for Go applications including 12-Factor apps.

## License 版权信息

本软件基于 [GPLv3](https://github.com/zouyonghe/Jinshuzhai-Bookstore/blob/main/LICENSE) 协议开源，使用者可以自由的使用、修改、再分发此软件源代码，但须遵守许可证相关条例。
