FROM golang:1.16-alpine


WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./

RUN go mod download
RUN go build -o /notfli

COPY entrypoint.sh /entrypoint.sh

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
