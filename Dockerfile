FROM golang:1.21-alpine as builder

ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o server

FROM scratch
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]

EXPOSE 5001