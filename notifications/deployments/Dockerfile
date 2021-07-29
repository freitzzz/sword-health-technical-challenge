# Based on:
# https://docs.docker.com/language/golang/build-images/
# https://www.redhat.com/sysadmin/tiny-containers
# https://chemidy.medium.com/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

# Plan:
# -> Pull alpine with go, copy source, build;
# -> Apply minimal image (scratch), copy built binary;
# -> Declare entrypoint

FROM golang:1.16-alpine AS builder

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY cmd/ ./cmd/

COPY internal/ ./internal/

RUN GOOS=linux CGO_ENABLED=0 go build -o /notifications ./cmd/notifications/main.go

FROM scratch

COPY --from=builder /notifications /notifications

CMD [ "/notifications" ]
