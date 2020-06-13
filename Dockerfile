FROM golang:1.14.4

LABEL maintainer "[golang app]"
WORKDIR /go/src

COPY . .

ENV GO111MODULE=on

RUN go mod download

RUN go get github.com/pilu/fresh

CMD ["fresh"]

