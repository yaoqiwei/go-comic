
FROM golang:rc-alpine3.13

ENV ENVIRONMENT test
RUN mkdir -p /go/src/fehu
WORKDIR /go/src/fehu
COPY go.mod /go/src/fehu/
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN go get -d github.com/swaggo/swag/cmd/swag
RUN go install github.com/swaggo/swag/cmd/swag
COPY . /go/src/fehu/
RUN swag init
RUN go build -o server.bin main.go

EXPOSE 20150
CMD /go/src/fehu/server.bin