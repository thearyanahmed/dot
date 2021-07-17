FROM alpine:3.6 as certificates

RUN apk add -U --no-cache ca-certificates

FROM golang:1.16-alpine3.14 as gobuilder

WORKDIR /go/src/github.com/thearyanahmed/dot

COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/dot

FROM scratch

COPY --from=gobuilder /go/bin/dot .

COPY --from=certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/dot"]