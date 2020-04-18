FROM golang:alpine AS builder
RUN apk add --no-cache git && mkdir -p $GOPATH/src/github.com/DTherHtun/generator && go get github.com/danielkvist/talk && go get gopkg.in/liderman/text-generator.v1
WORKDIR $GOPATH/src/github.com/DTherHtun/generator
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/generator .
FROM scratch
COPY --from=builder /go/bin/generator /go/bin/generator
ENTRYPOINT ["/go/bin/generator"]
