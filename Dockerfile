FROM golang:1.24-alpine

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o release-handler

EXPOSE 8080
CMD ["./release-handler", "server"]