FROM golang:latest

WORKDIR $GOPATH/src/github.com/yangyi422/CSForum
COPY . $GOPATH/src/github.com/yangyi422/CSForum
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./CSForum"]