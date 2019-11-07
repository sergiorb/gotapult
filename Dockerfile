FROM golang:latest AS builder
LABEL maintainer="Sergio Romero <s.romerobarra.tech@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/config.json .
ENTRYPOINT ["./main", "-configPath", "./config.json"]
