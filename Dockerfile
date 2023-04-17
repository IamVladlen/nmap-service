FROM golang:alpine as modules
COPY go.mod go.sum /go/nmap-service/modules/
WORKDIR /go/nmap-service/modules
RUN go mod download

FROM golang:alpine as builder
COPY --from=modules /go/pkg /go/pkg
WORKDIR /github.com/IamVladlen/nmap-service
COPY . .
RUN go build -o .bin/nmap/ ./cmd/app/main.go

FROM alpine as final
WORKDIR /nmap-service
COPY --from=builder /github.com/IamVladlen/nmap-service/.bin/nmap .
COPY --from=builder /github.com/IamVladlen/nmap-service/config config/
RUN apk add nmap --no-cache && rm -f /var/cache/apk/*
RUN apk add nmap-scripts
CMD [ "./main" ]