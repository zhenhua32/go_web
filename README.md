## 简介

使用 Gin 创建一个 web 项目.

主要参考的是掘金小册里的
[基于 Go 语言构建企业级的 RESTful API 服务](https://juejin.im/book/5b0778756fb9a07aa632301e).

## 版本

- 0.1.0 项目初始化
- v0.2.0 读取配置
- v0.3.0 记录日志
- v0.4.0 连接数据库
- v0.5.0 定义错误码
- v0.6.0 读取请求返回响应
- v0.7.0 添加核心逻辑, 用户数据的 CRUD
- v0.8.0 增加中间件
- v0.9.0 添加 JWT 认证
- v0.10.0 添加 HTTPS
- v0.11.0 添加 Makefile
- v0.12.0 添加版本信息
- v0.13.0 添加启动脚本
- v0.14.0 添加 Nginx 配置
- v0.15.0 添加测试的例子
- v0.16.0 添加 swagger 文档

## 运行

假设: 在项目根目录下运行命令

方式一: 在 docker 中运行 mysql, 本地启动服务器

```bash
# 后台启动 mysql 服务器
docker-compose up -d mysql
# 初始化数据库
docker-compose run --rm dbclient bash -c "cat /home/script/db.sql | mysql -hmysql -uroot -p1234"
# 运行服务器
go run ./
```

方式二: 在 docker 中运行 mysql, 本地编译二进制文件, 直接启动

```bash
# 后台启动 mysql 服务器
docker-compose up -d mysql
# 初始化数据库
docker-compose run --rm dbclient bash -c "cat /home/script/db.sql | mysql -hmysql -uroot -p1234"
# 编译, 应该会在当前目录下生成一个叫做 web 的二进制文件
make build
# 运行
web
```

方式三: 在 docker 中运行 mysql, 使用 systemd 接管服务

额外要求: 目录应该是 /home/go_web/, 否则需要更改配置中的路径

使用的配置路径是 /home/go_web/conf/config_abs.yaml

```bash
# 后台启动 mysql 服务器
docker-compose up -d mysql
# 初始化数据库
docker-compose run --rm dbclient bash -c "cat /home/script/db.sql | mysql -hmysql -uroot -p1234"
# 编译, 应该会在当前目录下生成一个叫做 web 的二进制文件
make build
# 复制文件到 systemd 的配置文件夹
cp conf/go_web.service /etc/systemd/system/
# 启动
systemctl start go_web
# 查看状态
systemctl status go_web
# 停止
systemctl stop go_web
```
