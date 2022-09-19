FROM golang:alpine as builder

WORKDIR /github.com/zhangrt/voyager1_platform
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="zhoujiajun@gsafety.com"

WORKDIR /github.com/zhangrt/voyager1_platform

COPY --from=0 /github.com/zhangrt/voyager1_platform/ ./
COPY --from=0 /github.com/zhangrt/voyager1_platform/resource/resource/rbac_model.conf ./resource/
COPY --from=0 /github.com/zhangrt/voyager1_platform/config.docker.yaml ./

EXPOSE 8888
ENTRYPOINT ./server -c config.docker.yaml
