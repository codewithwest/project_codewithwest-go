FROM golang:alpine3.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

EXPOSE 3071

USER nonroot:nonroot

CMD ["go","run","main.go"]
