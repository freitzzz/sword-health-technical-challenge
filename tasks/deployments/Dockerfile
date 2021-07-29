FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY cmd/ ./cmd/

COPY internal/ ./internal/

RUN GOOS=linux CGO_ENABLED=0 go build -o /tasks ./cmd/tasks/main.go

CMD [ "/tasks" ]