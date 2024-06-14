FROM golang:1.22.3-alpine3.19 AS base
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

FROM base AS build
RUN go build -ldflags "-s -w" -o /bin/server ./cmd/main.go

FROM alpine:3.19 AS prod
COPY --from=build /bin/server /bin/
EXPOSE 1080
CMD /bin/server >> /var/log/socks5.log 2>&1