# CurrencyExchange - 区块链货币兑换系统

一个基于 Vue 3 + Go 的去中心化货币兑换和资讯平台。

## 功能特性

- 💱 **货币兑换** - 支持多种加密货币之间的汇率兑换
- 📰 **资讯资讯** - 区块链新闻和资讯发布
- 👍 **互动功能** - 文章点赞功能
- 🔐 **用户认证** - 注册和登录系统
- 🌐 **RESTful API** - 完整的后端接口

## 技术栈

### 前端
- **Vue 3** - 渐进式 JavaScript 框架
- **TypeScript** - 类型安全的 JavaScript
- **Vite** - 新一代前端构建工具
- **Element Plus** - Vue 3 组件库
- **Pinia** - Vue 3 状态管理
- **Vue Router** - Vue 3 官方路由
- **Axios** - HTTP 客户端

### 后端
- **Go** - Google 开发的高性能编程语言
- **Gin** - Go HTTP Web 框架
- **GORM** - Go ORM 库

## 项目结构

```
CurrencyExchange/
├── frontend/              # 前端项目
│   ├── src/
│   │   ├── components/    # 组件
│   │   ├── router/        # 路由
│   │   ├── store/         # 状态管理
│   │   ├── types/         # TypeScript 类型
│   │   ├── views/         # 页面
│   │   ├── App.vue        # 根组件
│   │   └── main.ts        # 入口文件
│   ├── package.json       # 依赖配置
│   └── vite.config.ts     # Vite 配置
│
└── backend/               # 后端项目
    ├── config/            # 配置
    ├── controllers/       # 控制器
    ├── middlewares/       # 中间件
    ├── models/            # 数据模型
    ├── router/            # 路由
    ├── utils/             # 工具
    ├── go.mod             # Go 模块配置
    └── main.go            # 入口文件
```

## 快速开始

### 前置要求

- Node.js 18+
- Go 1.21+
- 数据库（MySQL/PostgreSQL/SQLite）

### 1. 克隆项目

```bash
git clone https://github.com/yourusername/CurrencyExchange.git
cd CurrencyExchange
```

### 2. 后端启动

```bash
cd backend

# 安装依赖
go mod download

# 配置数据库（编辑 config/config.yaml）
# ...

# 运行后端
go run main.go

# 或编译后运行
go build -o exchangeapp
./exchangeapp
```

后端默认运行在 `http://localhost:8080`

### 3. 前端启动

```bash
cd frontend

# 安装依赖
npm install

# 开发模式
npm run dev

# 构建生产版本
npm run build

# 预览生产版本
npm run preview
```

前端默认运行在 `http://localhost:5173`

## API 文档

### 认证

#### 注册
```
POST /api/auth/register
Content-Type: application/json

{
  "username": "user",
  "password": "password"
}
```

#### 登录
```
POST /api/auth/login
Content-Type: application/json

{
  "username": "user",
  "password": "password"
}
```

### 货币兑换

#### 获取汇率
```
GET /api/exchangeRates
```

#### 创建汇率（需要认证）
```
POST /api/exchangeRates
Authorization: Bearer <token>
Content-Type: application/json

{
  "fromcurrency": "BTC",
  "tocurrency": "ETH",
  "rate": 15.5
}
```

### 文章

#### 获取文章列表
```
GET /api/articles
```

#### 获取文章详情
```
GET /api/articles/:id
```

#### 创建文章（需要认证）
```
POST /api/articles
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "文章标题",
  "content": "文章内容",
  "author": "作者名"
}
```

#### 点赞文章（需要认证）
```
POST /api/articles/:id/like
Authorization: Bearer <token>
```

#### 获取文章点赞数
```
GET /api/articles/:id/like
```

## 页面说明

### 首页 (HomeView)
欢迎页面，提供导航到其他功能。

### 货币兑换 (CurrencyExchangeView)
货币兑换功能，支持选择源货币、目标货币和金额，显示兑换结果。

### 资讯列表 (NewsView)
显示所有文章/资讯列表。

### 资讯详情 (NewsDetailView)
显示单篇文章的详细内容和点赞数。

## 配置说明

### 后端配置 (backend/config/)

编辑配置文件设置数据库连接、端口等。

```yaml
app:
  host: ""
  port: "8080"

database:
  driver: "mysql"  # 或 postgres, sqlite
  host: "localhost"
  port: "3306"
  name: "currency_exchange"
  username: "root"
  password: "password"
```

### 前端配置 (frontend/src/axios.ts)

编辑 API 基础 URL：

```typescript
const axiosInstance = axios.create({
  baseURL: 'http://localhost:8080',  // 修改为你的后端地址
  timeout: 10000
});
```

## 开发计划

- [ ] 完善用户认证系统
- [ ] 添加用户资料页面
- [ ] 实现实时汇率更新
- [ ] 添加交易历史记录
- [ ] 支持更多加密货币
- [ ] 添加图表可视化
- [ ] 实现暗黑模式
- [ ] 多语言支持

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT

---

**开发者**: Clawdbot AI Agent
**最后更新**: 2026-01-31
