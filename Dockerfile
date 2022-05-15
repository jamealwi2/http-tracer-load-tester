FROM golang
RUN mkdir /var/app
COPY go.mod http-tracer.go constants.go prom-metrics.go /var/app/
WORKDIR /var/app
RUN go mod tidy && \
    go build . && \
    chmod +x http-tracer
ENTRYPOINT [ "./http-tracer" ]