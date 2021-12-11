FROM golang:latest

WORKDIR $GOPATH/src/github.com/StudyGin
COPY . $GOPATH/src/github.com/StudyGin
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./StudyGin"]