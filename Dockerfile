FROM alpine:3.19 as builder
RUN apk update && apk upgrade && \
      apk add go git
WORKDIR /irssi-dispatcher
COPY go.sum go.mod /irssi-dispatcher/
RUN go mod download
COPY *.go /irssi-dispatcher/
RUN go build

FROM alpine:3.19
COPY --from=builder /irssi-dispatcher/irssi-dispatcher /irssi-dispatcher/
ENTRYPOINT ["/irssi-dispatcher/irssi-dispatcher"]
