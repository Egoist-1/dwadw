.PHONY:  gen-api
# 数据库连接 URL
DB_URL := root:root@tcp(localhost:3306)/naming
TABLE_NAME := naming
BIG_TABLE_NAME := naming
# goctl 模版
TEMPLATE_PATH := C:/Users/PC/GolandProjects/landingpage/goctltempalate
# api 路由前缀
API_ROUTE_PREFIX :=
# api,proto文件所在目录
API_DIR_PATH := ./api
RPC_DIR_PATH := ./rpc
# 项目名
PROJECT_NAME := naming

# api,proto,swagger文件位置
API_PATH := $(API_DIR_PATH)/$(PROJECT_NAME).api
RPC_PATH := $(RPC_DIR_PATH)/$(PROJECT_NAME).proto
SWAGGER_PATH := $(API_DIR_PATH)/$(PROJECT_NAME).json

gen-api:	# 根据数据库表生成api文件
	api-gen api --table $(TABLE_NAME) --service $(PROJECT_NAME) --dir $(API_DIR_PATH) --url "$(DB_URL)" --prefix $(API_ROUTE_PREFIX)   -y
gen-rpc:	# 根据数据库表生成proto文件
	api-gen rpc --dir $(RPC_DIR_PATH) --url "$(DB_URL)" --table $(TABLE_NAME)
goctl-api:	# 根据api文件生成go代码
	goctl api go --api $(API_PATH) --dir ./api -style=goZero --home $(TEMPLATE_PATH)
goctl-rpc:	#根据rpc文件生成go代码
	goctl rpc protoc $(RPC_PATH) --go_out=./rpc --go-grpc_out=./rpc --zrpc_out=./rpc --home $(TEMPLATE_PATH) -m
model:	# 根据数据库生成model层代码
	goctl model mysql datasource --url="$(DB_URL)" --table="$(TABLE_NAME)" --dir="./rpc/internal/model/$(BIG_TABLE_NAME)" --cache --home $(TEMPLATE_PATH)
swagger:	# 生成swagger文档
	goctl api plugin -plugin "goctl-swagger swagger -filename $(SWAGGER_PATH) -host www.e-idear.com -basepath / -schemes http,https" -api $(API_PATH) -dir .
