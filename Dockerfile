FROM golang:1.21.0 AS builder

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod download

WORKDIR /src

COPY . /src

RUN go test ./... -v

RUN go vet

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o app

FROM debian as prod

RUN  echo "Asia/Shanghai" > /etc/timezone \
    && rm -f /etc/localtime \
    && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR /opt

COPY --from=builder /src/app /opt

EXPOSE 80

CMD ["./app"]

