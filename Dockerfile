#FROM golang:1.16-alpine
FROM alpine
WORKDIR /app

#COPY go.mod ./
#COPY go.sum ./
#COPY *.go ./
#RUN go mod download
#RUN go build -o /notfli

COPY notfli /notfli
ENV TZ=Asia/Shanghai
COPY entrypoint.sh /entrypoint.sh
RUN chmod 755 /notfli /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
CMD ["help"]

#./notfli
#
#-name
#
#-url
#
#-branch
#
#-user
#
#-time
#
#-result
