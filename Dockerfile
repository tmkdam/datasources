FROM golang:1.8.3-jessie

RUN apt-get update \
    && apt-get install -y cron \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /go/src/github.com/bloomapi/datasources
WORKDIR /go/src/github.com/bloomapi/datasources

COPY . /go/src/github.com/bloomapi/datasources
COPY docker-resources/crontab /etc/cron.d/runner
COPY docker-resources/runner.sh /runner.sh

RUN go get ./... || true
RUN go run configure.go
RUN make

CMD /runner.sh && cron -f
