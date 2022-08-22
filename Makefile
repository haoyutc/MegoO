default: build
all: build

# 构建镜像
build:
	go mod tidy && go mod download
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/app cmd/main.go
