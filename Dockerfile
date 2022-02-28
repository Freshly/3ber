FROM golang:rc-alpine3.15 as golang-build
WORKDIR /3ber
RUN apk update && apk add git

COPY go.mod go.sum .
RUN go mod download && go mod verify

COPY . .
RUN ./scripts/build.sh

FROM alpine:3.15.0
COPY --from=golang-build /3ber/bin/3ber /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/3ber"]
