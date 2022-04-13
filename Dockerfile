FROM golang:1.17-alpine as builder

RUN apk add --no-cache ca-certificates git

WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --ldflags '-s -w -extldflags "-static"' -tags netgo -installsuffix netgo -o ./redli


FROM alpine:3.9
RUN apk --no-cache add ca-certificates
COPY --from=builder /src/redli /redli
ENTRYPOINT ["/redli"]
