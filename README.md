# GoZeroAPI

基于 [go-zero](https://github.com/zeromicro/go-zero) 的 RESTful API 服务。  
实现了用户管理（查询、添加、删除、更新）等功能，支持 SQLite 数据库。

## 功能
- 用户信息增删改查
- SQLite 数据库存储
- go-zero REST API 框架
- 支持跨平台部署（Windows / Linux）

## 环境要求
- Go 1.20+
- SQLite3
- Git

## 安装与运行

### 1. 克隆项目
```bash
git clone https://github.com/yourname/GoZeroAPI.git
cd GoZeroAPI

### 2.安装依赖
go mod tidy

### 3. 运行（Linux / Mac）
CGO_ENABLED=1 go run main.go

### 运行（Windows）
go run main.go

### 4.API 示例
curl http://localhost:8888/users

### 5.项目结构
GoZeroAPI/
├── internal/     # 业务逻辑和处理
├── etc/          # 配置文件
├── user.api      # 接口文档
└── main.go       # 入口文件

许可证
MIT License