FROM golang:1.20.3-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM scratch
COPY --from=builder /app .
EXPOSE 8080 
CMD ["./main"]