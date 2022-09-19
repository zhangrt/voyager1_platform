# 项目介绍

## 代码结构

```shell
├── api
├── config
├── core
├── docs
├── global
├── initialize
│   └── internal
├── middleware
├── model
│   ├── request
│   └── response
├── plugin
├── repository
├── resource
├── router
├── service
├── source
└── utils

```

| 文件夹       | 说明                    | 描述                        |
| ------------ | ----------------------- | --------------------------- |
| `api`        | api层                   | api层 |
| `config`     | 配置包                  | config.yaml对应的配置结构体 |
| `core`       | 核心文件                | 核心组件(zap, viper, server)的初始化 |
| `docs`       | swagger文档目录         | swagger文档目录 |
| `global`     | 全局对象                | 全局对象 |
| `initialize` | 初始化 | router,redis,gorm,validator等的初始化 |
| `--internal` | 初始化内部函数 | gorm 的 longger 自定义,在此文件夹的函数只能由 `initialize` 层进行调用 |
| `middleware` | 中间件层 | 用于存放 `gin` 中间件代码 |
| `model`      | 模型层                  | 模型对应数据表              |
| `--request`  | 入参结构体              | 接收前端发送到后端的数据。  |
| `--response` | 出参结构体              | 返回给前端的数据结构体      |
| `plugin`     | 插件                  | 集成插件              |
| `repository` | repository层            | 存放数据处理逻辑 |
| `resource`   | 静态资源文件夹          | 负责存放静态文件                |
| `router`     | 路由层                  | 路由层 |
| `service`    | service层               | 存放业务逻辑问题 |
| `source` | source层 | 存放初始化数据的函数 |
| `utils`      | 工具包                  | 工具函数封装            |

## 开发约束

### Service mesh

* 远程调用配置通过结构体ServiceMesh来配置

* service name 根据业务规划再进行微服务解耦拆分后配置

* 通过resty包处理http请求

* 请求服务配置host、port为在服务网格中gateway的host、port值

![service mesh](resource\images\servicemesh.png)

### swagger文档

* `go install github.com/swaggo/swag/cmd/swag@latest`

* `swag init`

* `import _ "go_code/docs" => router.go`

* 通过 Addr/swagger/doc.json 地址可直接手动导入Apifox

swag init生成api文档条件，在api中以注释形式编写swagger文档

### ORM约束

* 数据库字段采用abc_efg

* Model采用AbcEfg形式

### JSON格式

* 首字母小写的驼峰命名方式

### Model格式

* field首字母大写的驼峰命名方式

### 文件格式

* 小写字母，当有多个单词时通过 _ 连接

### git commit message 约束

**_type: subject(scope)_**

* type：用于说明commit的类别，规定为如下几种
* feat：新增功能；
* fix：修复bug；
* docs：修改文档；
* refactor：代码重构，未新增任何功能和修复任何bug；
* build：改变构建流程，新增依赖库、工具等（例如webpack修改）；
* style：仅仅修改了空格、缩进等，不改变代码逻辑；
* perf：改善性能和体现的修改；
* chore：非src和test的修改；
* test：测试用例的修改；
* ci：自动化流程配置修改；
* revert：回滚到上一个版本；
* scope：【可选】用于说明commit的影响范围
* subject：commit的简要说明，尽量简短

## 代码质量检查

### pre-commit

pip install pre-commit，需要安装python环境

#### Doc

<https://pre-commit.com/>

#### Run

* pre-commit install
* pre-commit run --all-files

* commit
![service mesh](resource\images\hook-commit.png)

* commit-msg
![service mesh](resource\images\hook-commit-msg.png)

* push
![service mesh](resource\images\hook-push.png)

### golangci-lint

#### 安装

`go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0`

#### 基础使用

在项目根路径执行 `golangci-lint run`

#### 更多

<https://github.com/golangci/golangci-lint>

## 单元测试

在*_test.go文件中有三种类型的函数：测试函数、基准测试、示例函数

* 功能测试：函数名必须以Test开头，函数参数必须是*testing.T。测试成立逻辑行为是否正确。
* 性能测试：函数名必须以Benchmark开头，函数参数必须是*testing.B。测试函数的性能。
* 示例测试：函数名必须以Example开头，函数参数无要求。为文档提供示例文档。
* 测试函数中某条测试用例执行结果与预期不符可调用t.Error()或t.Errorf()方法记录日志并标记测试失败

测试用例具有四种形式

* 基本测试用例：TestXxx(t *testing.T)
* 压力测试用例：BenchmarkXxx(b *testing.B)
* 测试控制台输出：Example_Xxx()
* 测试主函数：TestMain(m *testing.M)

### Test file

* 测试文件名字必须是xxx_test后缀
* 测试函数必须是Test开头（否则报错，no test to run）

### 单元测试覆盖率

* go test ./... -coverprofile=sonar/cov.out

* sonar-scanner.bat -D"sonar.projectKey=test" -D"sonar.sources=." -D"sonar.host.url=<http://192.168.244.142:9100>" -D"sonar.login=ba6f1edcbf7aa20bf01c02306413f2947bc180ee"

## Websocket

<https://github.com/gorilla/websocket>

## 学习扩展

## 关于Gin学习推荐

<https://github.com/skyhee/gin-doc-cn.git>

## 关于GO学习推荐

<https://github.com/golang101/golang101>

## 【腾讯文档】智慧园区Demo服务端框架介绍

<https://docs.qq.com/doc/DVURVSWlDS01Tc29I>

## 【腾讯文档】服务网格Istio

<https://docs.qq.com/doc/DVUVaRGVkWEFjanpi>
