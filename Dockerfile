FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /supportportal ./cmd/api

FROM alpine:3.22

RUN addgroup -S supportportal && adduser -S supportportal -G supportportal

WORKDIR /app
COPY --from=builder /supportportal /app/supportportal

USER supportportal

EXPOSE 8080

ENTRYPOINT ["/app/supportportal"]
