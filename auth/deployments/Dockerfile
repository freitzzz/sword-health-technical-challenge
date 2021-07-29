FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY api/ ./api/

COPY cmd/ ./cmd/

COPY internal/ ./internal/

RUN go build -o /auth ./cmd/auth/main.go

CMD [ "/auth" ]