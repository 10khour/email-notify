FROM golang:1.16.4-buster as builder
RUN mkdir /app/
ADD . /app/
WORKDIR /app/
ENV GOPROXY=https://goproxy.cn CGO_ENABLED=0
RUN go build

FROM centos:7
COPY --from=builder /app/notify /bin/notify

