FROM alpine:3.12

COPY .target/bin/prometheus-custom-metrics /bin

CMD /bin/prometheus-custom-metrics
