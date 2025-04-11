FROM golang:1.24-alpine3.21 AS base
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

FROM base AS build
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static" -s -w' -o /server ./cmd/main.go

FROM gcr.io/distroless/static-debian11
COPY --from=build /server /server
USER nonroot:nonroot
WORKDIR /
CMD ["/server"]
