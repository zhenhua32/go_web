BASEDIR = $(shell pwd)

# build with verison infos
versionDir = "tzh.com/web/pkg/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=UTC date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"

all: gotool build
build: updoc
	go build -ldflags ${ldflags} ./
run:
	go run -ldflags ${ldflags} ./
clean:
	rm -f web
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}
gotool:
	go fmt ./
	go vet ./
ca:
	MSYS_NO_PATHCONV=1 openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=CN/ST=SH/L=SH/O=CoolCat/OU=CoolCat Software/CN=127.0.0.1/emailAddress=coolcat@qq.com"
mysql:
	docker-compose up -d mysql
dbcli:
	docker-compose run --rm dbclient
updoc:
	swag init

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
	@echo "make ca - 生成证书文件"
	@echo "make mysql - 启动 mysql 服务器"
	@echo "make dbcli - 连接到 mysql 命令行"

.PHONY: run clean gotool ca mysql dbcli help