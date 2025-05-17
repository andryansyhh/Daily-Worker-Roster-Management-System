FROM golang:1.21-alpine

RUN apk add --no-cache gcc musl-dev sqlite

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -o main main.go && ls -al /app

EXPOSE 8089

CMD ["/app/main"]