# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.9-alpine3.6

# Copy the local package files to the container's workspace.
#ADD ./server.json /go/bin

RUN apk add --update git && \
    go get -t github.com/yinqiwen/gsnova/remote/server  && \
    go install github.com/yinqiwen/gsnova/remote/server 

#WORKDIR /go/bin
# Run the outyet command by default when the container starts.
#ENTRYPOINT ["/go/bin/vps"]
CMD ["/go/bin/server","-key", "809240d3a021449f6e67aa73221d42df942a308a", "-tcp", ":9443", "-quic", ":9443", "-http", ":9444", "-kcp", ":9444", "-http2", ":9445", "-tls", ":9446"]
# Document that the service listens on port 9443.
EXPOSE 9443 9445 9446