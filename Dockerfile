FROM golang:1.24-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download;  \
    go mod tidy

COPY . .

EXPOSE 3071

CMD ["air"]
