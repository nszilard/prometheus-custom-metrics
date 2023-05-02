FROM golang:alpine AS builder
WORKDIR /go/prometheus-custom-metrics
COPY . .
RUN go mod download -json
RUN go build -o /bin/prometheus-custom-metrics

FROM scratch
COPY --from=builder /bin/prometheus-custom-metrics /bin/prometheus-custom-metrics
ENTRYPOINT ["/bin/prometheus-custom-metrics"]
