FROM golang:1.17-alpine as base

RUN mkdir /app
WORKDIR /api
COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["go", "run", "main.go"]