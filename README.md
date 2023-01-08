# blog
## 使用说明
### 1.1环境要求
```
-golang版本 >= v1.18
-mysql >= 5.0
-redis >= 5.0.14
```

### 1.2运行项目
```
# 克隆项目
git clone git@github.com:fishworm96/blog.git
```
```
# 项目配置
必须先配置再启动
根目录下./conf/config.yaml文件
name: "blog"  // 项目名称
mode: "dev"   // 模式
port: 8080  // 运行端口
version: "v0.1.4" // 当前版本
start_time: "2022-07-01" // 开始时间
machine_id: 1 // 雪花算法创建新节点值
email_secret_code: "" // 可选：email smtp秘钥
access_key: "" // 可选配置项：oss access key
secret_key: "" // 可选配置项：oss secret key
bucket: "" // 可选配置项：oss 桶名称
img_url: "" // 可选配置项：解析域名地址

auth:
  jwt_expire: 8760 // jwt 过期时间
log:
  level: "info" // 日志等级
  filename: "blog.log"  // 日志文件名称
  max_size: 200 // 最大大小
  max_age: 30 // 最长时间
  max_backups: 7 // 最多文件
mysql:
  host: "127.0.0.1"  // 重要：mysql host
  port: 3306 // 重要：端口
  user: "root" // 重要：账号
  password: "root" // 重要：密码
  dbname: "blog" // 重要：数据库名称
  max_open_conns: 200 // 重要：最多连接数
  max_idle_conns: 50 // 重要：最大空闲数
redis:
  host: "127.0.0.1" // 重要：redis host
  port: 6379 // 重要：端口
  password: "" // 重要：密码
  db: 0 // 重要：默认数据库
  pool_size: 100 // 重要：最大连接池
```

### 1.3启动项目
```
# 手动启动
go run main.go ./conf/config.yaml
```
```
# 使用热更新
go install github.com/cosmtrek/air@latest
在终端中使用
air

# 配置air
根目录下.air.conf
```