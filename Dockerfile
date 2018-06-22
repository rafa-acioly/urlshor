FROM golang as builder

ADD . /go/src/github.com/rafa-acioly/urlshor
WORKDIR /go/src/github.com/rafa-acioly/urlshor
RUN curl -s https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure && CGO_ENABLED=0 go build -ldflags '-s -w'

FROM scratch
WORKDIR /
COPY --from=builder /go/src/github.com/rafa-acioly/urlshor/urlshor /
ADD ./static /go/src/github.com/rafa-acioly/urlshor/static

CMD ["/urlshor"]
