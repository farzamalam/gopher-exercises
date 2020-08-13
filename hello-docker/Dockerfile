FROM golang:1.12

WORKDIR /go/src/hello-docker

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["hello-docker"]