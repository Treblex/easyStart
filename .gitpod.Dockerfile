FROM alpine

RUN apk update
RUN apk add --no-cache go
# Configure Go
ENV GOPATH /workspace/go
ENV PATH $GOPATH/bin:$PATH
# ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

COPY  ./server/config/zoneinfo.zip /opt/zoneinfo.zip
ENV ZONEINFO /opt/zoneinfo.zip
WORKDIR /workspace/go-echo-demo/server
EXPOSE 8080

