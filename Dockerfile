FROM golang:1.22.1-alpine3.18 AS builder
WORKDIR /app
ADD . /app/
RUN go mod download && go mod verify
RUN go build -ldflags "-s -w"

FROM alpine:3.18 AS runner
RUN apk add ffmpeg
COPY --from=builder /app/ffserver /usr/bin/ffserver

WORKDIR /tmp
EXPOSE 3000
CMD ["ffserver"]