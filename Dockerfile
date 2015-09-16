FROM golang:latest
MAINTAINER Kiet Thanh Vo <kiet@ventuso.net>
# Get godep
RUN go get github.com/tools/godep
USER root

RUN mkdir /app  \
    mkdir /app/ws-chat
ADD . /app/ws-chat

#WORKDIR /app/ws-chat
#RUN godep restore
#RUN go build
#EXPOSE 8080