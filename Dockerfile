FROM golang:latest

WORKDIR $GOPATH/src/github.com/yangyi422/CSForum
COPY . $GOPATH/src/github.com/yangyi422/CSForum
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./CSForum"]