# Redis availability checker

## Install dependencies

```bash
go mod vendor
```

## Run in single-server mod

```bash
go run main.go --server=127.0.0.1:6379
```

## Run in cluster mod

```bash
go run main.go --cluster=redis1:6379,redis2:6379,redis3:6379
```
