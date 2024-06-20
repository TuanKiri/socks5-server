
FROM golang:1.22.3-alpine3.19 AS base
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

FROM base AS build
RUN go build -ldflags "-s -w" -o /bin/server ./cmd/main.go

FROM alpine:3.19 AS prod
COPY --from=build /bin/server /bin/
ARG USER_UID
ARG USER_GID
ARG USERNAME=user
RUN apk update && apk add --no-cache \
    sudo \
    shadow \
    && addgroup -g $USER_GID $USERNAME \
    && adduser -u $USER_UID -G $USERNAME -s /bin/sh -D $USERNAME \
    && mkdir -p /var/log/socks5 \
    && chown $USER_UID:$USER_GID /var/log/socks5
USER $USERNAME
EXPOSE 1080
CMD /bin/server >> /var/log/socks5/socks5.log 2>&1