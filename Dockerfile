FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache nodejs npm
RUN npm install -g nodemon

COPY go.mod go.sum ./
RUN go mod download

COPY . .

  #go build -o main ./main.go

EXPOSE 3072

# Corrected CMD instruction:
CMD ["sh", "-c", "while true; do go run main.go; sleep 1; done"]
