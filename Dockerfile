FROM golang:1.22 AS base
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOPRIVATE="github.com/seyio91"

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./
COPY pkg ./pkg/

RUN go build -o /main

RUN wget https://github.com/FairwindsOps/pluto/releases/download/v5.19.0/pluto_5.19.0_linux_amd64.tar.gz

RUN tar -xzf pluto_5.19.0_linux_amd64.tar.gz -C /usr/bin/

EXPOSE 2112

CMD ["/main"]
