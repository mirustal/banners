FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main ./cmd/main.go

FROM alpine:latest 

WORKDIR /root/

COPY --from=builder /app/main .
COPY app.env .

RUN chmod +x main

EXPOSE 8081

CMD ["./main"]
