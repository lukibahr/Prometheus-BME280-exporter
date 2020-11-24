FROM golang:1.15-alpine as builder
WORKDIR /go/src/app
COPY . /go/src/app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/exporter .

FROM gcr.io/distroless/base@sha256:be45bda793f6c5798ebca8f9ccfde7d312ba386858f27d5a24084fbe48db9d3c
COPY --from=builder /go/bin/exporter /
ENTRYPOINT ["/exporter"]
CMD [""]
