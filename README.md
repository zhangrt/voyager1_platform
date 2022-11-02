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

* `swag init --parseDependency --parseInternal 项目引入了基础组件需要识别外部依赖`

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

* go test ./... -coverprofile="sonar/cov.out"

* sonar-scanner.bat -D"sonar.projectKey=test" -D"sonar.sources=." -D"sonar.host.url=<http://192.168.244.142:9100>" -D"sonar.login=ba6f1edcbf7aa20bf01c02306413f2947bc180ee"

## Websocket

<https://github.com/gorilla/websocket>

## 基础组件Core

### 初始化配置

Core.New().Config().Config.....

### import相关组件 如Auth Cache 等

* platform服务导入luna相关包 import "github.com/zhangrt/voyager1_core/auth/luna"

* 业务服务导入star相关包 import "github.com/zhangrt/voyager1_core/auth/star"

* 导入Cache相关包 import "github.com/zhangrt/voyager1_core/cache"

* 以上诸如此类

#### 配置解析 application.yaml

* **system**

  系统初始化的一些配置：mod运行模式、role服务角色等等

  默认情况下：mod为test时，访问将没有任何auth拦截；mode为develop时，访问将只验证token合法性，而不查验用户角色权限

  role的定义: platform服务role为：luna；业务服务role为：star

  service-casbin-name和service-jwt-name为实现了ICasbin和IJwt接口实现bean的名称，用于查找接口实现的bean
  
```yaml
system:
  # test、develop、release & develop skip authenticate & test skip all auth
  mode: release
  # luna - platform/ star - business
  role: luna
  # 公开的路由配置
  public-routes:
    - /**/login
    - /**/register
    - /**/logout
    - /**/person/departments/accountorphoneoremail/*
  # 多点登录拦截 false - 单点登录 / true - 多点登录
  use-multipoint: true
  # grpc集成方式 => springboot/java
  use-grpc-type: java
  # auth 权限接口的实现类bean的名称
  service-casbin-name: casbinService
  service-jwt-name: jwtService
```

* **cors**

```yaml
# 跨域配置
# 需要配合 过滤器或拦截器 使用
cors:
  mode: allow-all # 放行模式: allow-all, 放行全部; whitelist, 白名单模式, 来自白名单内域名的请求添加 cors 头; strict-whitelist 严格白名单模式, 白名单外的请求一律拒绝
  whitelist:
    - allow-origin: localhost
      allow-headers: content-type
      allow-methods: GET, POST, DELETE, PUT
      expose-headers: Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type
      allow-credentials: true # 布尔值
    - allow-origin: 127.0.0.1
      allow-headers: content-type
      allow-methods: GET, POST, DELETE, PUT
      expose-headers: Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type
      allow-credentials: true # 布尔值
```

* **casbin**

```yaml
# casbin configuration 集成 springboot 的配资
#自版本0.0.11开始, casbin-spring-boot-starter 默认为数据库表结构添加ID字段
#0.0.11之前的版本升级时, 需要用户手动添加ID字段
casbin:
  #是否开启Casbin,默认开启
  enableCasbin: true
  #是否使用线程同步的Enforcer,默认false
  useSyncedEnforcer: false
  #是否开启策略自动保存，如适配器支持该功能，默认开启
  autoSave: true
  #存储类型[file,jdbc]，目前支持的jdbc数据库[mysql(mariadb),h2,oracle,postgresql,db2]
  #欢迎编写并提交您所使用的jdbc适配器，参见：org.casbin.adapter.OracleAdapter
  #jdbc适配器将主动寻找您在spring.datasource配置的数据源信息
  #默认使用jdbc,并使用内置h2数据库进行内存存储
  storeType: jdbc
  #当使用jdbc时,定制化数据库表名,默认表名是casbin_rule
  tableName: casbin_rule
  #数据源初始化策略[create(自动创建数据表,如已创建则不再进行初始化),never(始终不进行初始化)]
  initializeSchema: create
  #本地模型配置文件地址,约定默认读取位置:classpath:casbin/model.conf
  model: casbin/rbac_model.conf
  #如默认位置未找到模型配置文件,且casbin.model未正确设置,则使用内置默认rbac模型,默认生效
  useDefaultModelIfModelNotSetting: true
  #本地策略配置文件地址,约定默认读取位置:classpath:casbin/policy.csv
  #如默认位置未找到配置文件，将会抛出异常
  #该配置项仅在casbin.storeType设定为file时生效
  policy:
  #是否开启CasbinWatcher机制,默认不开启
  #如开启该机制,则casbin.storeType必须为jdbc,否则该配置无效
  enableWatcher: false
  #CasbinWatcher通知方式,默认使用Redis进行通知同步,暂时仅支持Redis
  #开启Watcher后需手动添加spring-boot-starter-data-redis依赖
  watcherType: redis
  #异常抛出时机控制
  exception:
    #删除策略失败时是否抛出异常
    removePolicyFailed: false
```

* **jwt**

    定义Token的一些基本配置

```yaml
# jwt configuration
jwt:
  # 加密key
  signing-key: gsafety-auth-key-1024
  # 测试new token 可以用较小的过期时间
  # 单位 秒
  expires-time: 604800 # 过期时间7天
  buffer-time: 86400 # 缓存过期时间1天，在6-7天之间会用old token 生成 new token
  issuer: gsafety
```

* **auth-key**

    auth用到的一些key name的定义，例如：request header里面的token key名称可以定义为：x-token（前后端一致即可）

```yaml
# 定义auth用到的一些key值,用于前后端校验时获取请求响应中的token等关键信息
# key名称的配置由项目自定义
auth-key:
  token: x-token                     # token key的命名
  expires-at: expiresAt              # token过期 key的命名
  refresh-token: new-token           # new token key的命名
  refresh-expires-at: new-expires-at # new token 过期时间 key的命名
  user: claims                       # 用户信息 key的命名
  user-id: x-user-id                 # 用户id key的命名
  reload: reload                     # 客户端是否需要重载token的配置，作为response的header中key名称的配置
```

* **grpc**

    gRPC的一些配置，与springboot集成的配置和普通java集成的配置，两种配置方式
 
```yaml
# grpc 配置信息
grpc:
# gRPC自定义配置与golang版本一致
  j-server:
    network: tcp
    host: 127.0.0.1
    port: 2568
  j-client:
    host: 127.0.0.1
    port: 2568
# 兼容 springboot 集成 gRPC的配置，若采用springboot集成，必不可少
  server:
    port: 2567
    address: 127.0.0.1
  client:
    # 定义一个gRpc的连接信息auth-service，与注解@GrpcClient(name)的name保持一致
    auth-service:
      negotiationType: PLAINTEXT
      address: static://127.0.0.1:2567
```

* **service-mesh**

    服务网格配置，定义服务的基本信息，用于基本的服务发现等

```yaml
# service-mesh 核心配置 host & port
service-mesh:
  namespaces:
    - name: default
      loadbalance: 0.0.0.0
      services:
        - name: facility-service
          host: facility
          # v0.1 subset版本预留配置,现阶段未用上
          subset: v1
          port: 9998
          prefix: /gsafety
        - name: test-service
          host: test
          subset: v1
          port: 9998
          prefix: /test
    - name: gsafety-test
      loadbalance: 0.0.0.0
      services:
        - name: facility-service
          host: facility
          # v0.1 subset版本预留配置,现阶段未用上
          subset: v1
          port: 9998
          prefix: /gsafety
        - name: test-service
          host: test
          subset: v1
          port: 9998
          prefix: /test
```

## 学习扩展

## 关于Gin学习推荐

<https://github.com/skyhee/gin-doc-cn.git>

## 关于GO学习推荐

<https://github.com/golang101/golang101>

## 【腾讯文档】智慧园区Demo服务端框架介绍

<https://docs.qq.com/doc/DVURVSWlDS01Tc29I>

## 【腾讯文档】服务网格Istio

<https://docs.qq.com/doc/DVUVaRGVkWEFjanpi>
