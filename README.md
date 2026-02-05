# golang-backend

基于 Gin 框架的 Go 后端项目

## 环境要求

- Go 1.25.6+
- MySQL
- Redis

## 快速开始

```bash
# 安装依赖
go mod tidy

# 运行项目 (设置环境变量)
GO_ENV=local go run main.go
```

## 项目结构

```text
golang-backend/
├── cmd/                    # 命令行工具
├── configs/                # 配置文件
│   ├── dev.yaml           # 开发环境配置
│   ├── local.yaml         # 本地环境配置
│   ├── prod.yaml          # 生产环境配置
│   └── test.yaml          # 测试环境配置
├── internal/               # 内部模块
│   ├── controllers/       # 控制器层，处理请求和响应
│   ├── converter/         # 数据转换器
│   ├── core/              # 核心模块（加密、会话等）
│   ├── dao/               # 数据访问层
│   ├── entity/            # 实体定义（请求/响应结构）
│   ├── infra/             # 基础设施层（MySQL、Redis初始化）
│   ├── middleware/        # 中间件（认证、日志等）
│   ├── models/            # 数据库模型
│   ├── routers/           # 路由配置
│   ├── services/          # 服务层，核心业务逻辑
│   ├── setting/           # 配置加载
│   └── utils/             # 内部工具包
│       ├── ctxhelper/     # 上下文助手
│       ├── logger/        # 日志工具
│       └── resp/          # 响应工具
├── pkg/                    # 公共工具包
│   ├── bcrypt/            # 密码加密
│   ├── constants/         # 常量定义
│   ├── db/                # 数据库连接
│   ├── logger/            # 日志组件
│   ├── redis/             # Redis 连接
│   ├── snowflake/         # 雪花算法 ID 生成
│   ├── xjwt/              # JWT 工具
│   └── xsignature/        # 签名工具
├── scripts/                # 脚本文件
├── main.go                 # 入口文件
├── go.mod                  # Go 模块定义
└── go.sum                  # 依赖版本锁定
```

## 请求流程

```
前端请求 → 中间件 → 控制器 → 服务层 → 数据访问层 → 数据库
                                                    ↓
前端响应 ← 控制器(entity) ←────────────────────────←┘
```
