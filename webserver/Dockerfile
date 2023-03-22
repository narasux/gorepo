# -------------- builder container --------------
# Build the manager binary
FROM golang:1.19 as builder

WORKDIR /go/src/

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o gowebserver ./main.go

# -------------- runner container --------------
FROM alpine:3.15 AS runner

RUN apk --update --no-cache add bash

WORKDIR /data/workspace

RUN mkdir -p /data/workspace/assets

ENV ASSETS_DIR=/data/workspace/assets

COPY --from=builder /go/src/gowebserver /usr/bin/gowebserver

ENTRYPOINT ["gowebserver"]
