# Jinshuzhai-Bookstore

An online bookstore system based on Gin.

## Introduction 介绍

一个基于Gin、gorm、viper、zap等库的web服务器，实现了网上书店相关接口。

### Summary 概要

### Features 特性

## Requirements 必要条件

参考go.mod文件。

## Configuration 配置

在conf/config.yaml文件中修改配置。

```yaml
runmode: debug               # 开发模式, debug, release, test
addr: :8080                  # HTTP绑定端口
name: bookstore-server       # API Server的名字
url: http://127.0.0.1:8080   # pingServer函数请求的API服务器的ip:port
max_ping_count: 10           # pingServer函数try的次数
jwt_secret: Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5
tls:
  addr: :8081                # tls端口
  cert: conf/server.crt      # cert文件
  key: conf/server.key       # key文件
log:
  log_level: DEBUG           # 日志级别
  log_file: blog/server.log  # 日志位置
  max_backups: 5             # 日志最大备份数量
  max_size: 1                # 日志最大文件大小，单位MB
  max_age: 30                # 日志最大保留天数
  compress: false            # 是否压缩
db:
  name: bookstore_server     # 数据库名称
  addr: 127.0.0.1:3306       # 数据库所在IP及端口号
  username: root             # 数据库登陆用户名
  password: aa11bb22cc33     # 数据库登陆密码
docker_db:
  name: db_apiserver
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
#启动服务

./server.sh start

#终止服务

./server.sh stop
```

## Development 开发文档

## Changelog 更新日志

| 日期        | 事项                 | 开发者    |
|:---------:|:------------------:|:------:|
| 2022-4-14 | 初始化项目，完成配置和日志记录代码  | buding |
| 2022-4-15 | 完成版本信息，调整配置和日志记录代码 | buding |

## FAQ - 常见问题（常见问题。）

## Support 支持

如果需要反馈问题，可以直接提交 [Issues](https://github.com/zouyonghe/Jinshuzhai-Bookstore/issues) 或联系邮箱。

### Dos 文档

### Contact 联系

name: buding

email: 1259085392z@gamil.com

## Authors and acknowledgment 贡献者和感谢

The [Go](https://github.com/golang/go) programming language

[Gin](https://github.com/gin-gonic/gin) is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin

[Gorm](https://github.com/go-gorm/gorm)  is the fantastic ORM library for Golang, aims to be developer friendly

## License 版权信息

本软件基于 [GPLv3](https://github.com/zouyonghe/Jinshuzhai-Bookstore/blob/main/LICENSE) 协议开源，开发者可以自由的使用、修改、再分发此软件源代码，但须遵守相关协议。
