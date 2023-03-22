# -------------- builder container --------------
FROM golang:1.19 as builder

WORKDIR /go/src/

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o gohello ./main.go

# -------------- runner container --------------
FROM alpine:3.15 AS runner

RUN apk --update --no-cache add bash

WORKDIR /data/workspace

COPY --from=builder /go/src/gohello /usr/bin/gohello

ENTRYPOINT ["gohello"]
