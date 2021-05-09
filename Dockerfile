FROM golang:1.15-alpine as builder
WORKDIR /go/src/app
COPY . /go/src/app
RUN CGO_ENABLED=0 GOARCH=arm GOARM=7 GOOS=linux go build -a -installsuffix cgo -o /go/bin/exporter .

FROM alpine:3.13
COPY --from=builder /go/bin/exporter /
ENTRYPOINT ["/exporter"]
CMD [""]
