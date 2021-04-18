FROM golang:1.16-alpine
WORKDIR /dcard-ratelimit-middleware
ADD . /dcard-ratelimit-middleware
RUN cd /dcard-ratelimit-middleware && go build
ENTRYPOINT ["./dcard-ratelimit-middleware"]
