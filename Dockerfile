FROM golang:1.16
WORKDIR $GOPATH/src/github.com/bredbrains/tthk-wish-list
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8080
CMD ["tthk-wish-list"]