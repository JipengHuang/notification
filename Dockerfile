#FROM golang:1.16-alpine
FROM alpine
WORKDIR /app

#COPY go.mod ./
#COPY go.sum ./
#COPY *.go ./
#RUN go mod download
#RUN go build -o /notfli

COPY notfli /notfli
COPY entrypoint.sh /entrypoint.sh

RUN apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    chmod 755 /notfli /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
CMD ["help"]

