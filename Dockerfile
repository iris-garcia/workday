FROM golang:1.13

RUN mkdir -p /go/src/github.com/iris-garcia/workday
WORKDIR /go/src/github.com/iris-garcia/workday

COPY . /go/src/github.com/iris-garcia/workday
RUN go get github.com/magefile/mage
RUN mage build

USER nobody

CMD ["./api_server"]
