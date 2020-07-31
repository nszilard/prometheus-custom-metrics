FROM alpine:3.12

COPY ./prometheus-custom-metrics /bin

ENV APPLICATION_ENV="production"

EXPOSE 8080
CMD /bin/prometheus-custom-metrics
