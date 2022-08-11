FROM golang:1.16-alpine

WORKDIR /src/go_wrapper

COPY go_wrapper/. /src/go_wrapper/
RUN go get ./... && rm -rf /src/go_wrapper/handler

CMD go_wrapper < /dev/null
LABEL racetrack-component="fatman"
