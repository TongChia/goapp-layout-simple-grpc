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
- `api/gen`
- `cmd`
- `configs`
- `internal/biz`
- `internal/data`
- `internal/server`
- `internal/service`