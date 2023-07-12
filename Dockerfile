FROM golang:1.17-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/main .

CMD ["./main"]
