FROM golang:latest

# Get godep
RUN go get github.com/tools/godep
USER root
RUN mkdir /app
ADD . /app/ws-chat
RUN cd /app/ws-chat &&\
    godep restore   &&\
	go build
 
EXPOSE 8080