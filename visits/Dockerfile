FROM golang:1.13


WORKDIR /go/src/visits
COPY ./go.mod .
RUN go get -d -v ./...

COPY . .

RUN go install .

CMD ["visits"]
