FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o iotea .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates samba-client

WORKDIR /root/
COPY --from=builder /app/iotea .

EXPOSE 8080

CMD ["./iotea"]