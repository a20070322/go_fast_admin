SHELL=/usr/bin/env bash
uname=${shell uname}
appName=fast_admin
build_path= ./dist/
server_app_path = /app/fast_admin/
server_port = 22
server_user = root
server_host = 101.42.96.67


start:
	@env go run main.go start
auto:
	@env go run main.go AutoCurd -c .tmp/config.json

ent_m:
	@env go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
	@env go run main.go ent

linux:
	@echo 编译linux版本文件
	@env CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o ${build_path}${appName}_linux

# 项目服务器初始化
deploy_init:
	@echo "部署初始化文件...."
	@scp -P ${server_port} -r ./config/config.yaml ${server_user}@${server_host}:${server_app_path}

# 部署至配置服务器
deploy:
	@echo "部署项目...."
	@ssh -p ${server_port} ${server_user}@${server_host} "systemctl stop ${appName}"
	@scp -P ${server_port} -r ./dist/${appName}_linux ${server_user}@${server_host}:${server_app_path}
	@ssh -p ${server_port} ${server_user}@${server_host} "systemctl start ${appName}"
	@echo "部署成功"
#deploy_service:
#	@echo "部署服务文件...."
#	@scp -P 6079 -r ./${linux_app_name}.service ${linux_scp_user}@${linux_scp_host}:/etc/systemd/system/
#
