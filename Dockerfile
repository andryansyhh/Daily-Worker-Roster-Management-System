FROM golang:1.21-alpine

# 1. Install dependencies
RUN apk add --no-cache gcc musl-dev sqlite

# 2. Set working directory
WORKDIR /app

# 3. Copy go.mod & go.sum and download modules
COPY go.mod go.sum ./
RUN go mod download

# 4. Copy all files
COPY . .

# 5. Debug check: is main.go actually there?
RUN ls -al /app && cat /app/main.go

# 6. Build binary (verbose mode)
RUN go build -v -o main main.go

# 7. Confirm binary exists
RUN ls -al /app

# 8. Expose & run
EXPOSE 8089
CMD ["./main"]
