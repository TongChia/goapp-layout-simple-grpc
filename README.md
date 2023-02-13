goapp-layout-template
===

## Install built tools
```sh
make init
```

## Generate API
```sh
make api
```

## Build
```sh
make build
```

## Directory Structure
```
.
├── api
│   ├── buf.gen.yaml
│   ├── buf.lock
│   ├── buf.yaml
│   ├── gen
│   └── helloworld
├── cmd
│   └── main.go
├── configs
│   └── config.yaml
└── internal
    ├── biz
    ├── conf
    ├── data
    ├── server
    └── service
```

- `api`

    gRPC Protocol 模式文件，协议定义文件。

- `cmd`

    每个应用程序的目录名应该与你想要的可执行文件的名称相匹配(例如，/cmd/myapp)。
    
    不要在这个目录中放置太多代码。如果你认为代码可以导入并在其他项目中使用，那么它应该位于 /pkg 目录中。如果代码不是可重用的，或者你不希望其他人重用它，请将该代码放到 /internal 目录中。你会惊讶于别人会怎么做，所以要明确你的意图!
    
    通常有一个小的 main 函数，从 /internal 和 /pkg 目录导入和调用代码，除此之外没有别的东西。

- `configs`
    
    配置文件模板或默认配置。

    将你的 confd 或 consul-template 模板文件放在这里。

- `internal`
    
    私有应用程序和库代码。这是你不希望其他人在其应用程序或库中导入代码。请注意，这个布局模式是由 Go 编译器本身执行的。有关更多细节，请参阅Go 1.4 release notes 。

- `internal/server`

    gRPC/HTTP 服务相关代码, 包括注册中间件和配置证书等.

- `internal/data`

    数据模型和`repository`, 仅包含必要的增删改查, 不涉及具体的业务逻辑.

- `internal/biz`
    
    (business) 业务逻辑代码, 为`service`提供处理业务的接口.

- `internal/service`

    (API entry) 客户端请求的入口, 处理或校验参数并调用处理业务接口.
