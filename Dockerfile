FROM golang as builder

ADD . /go/src/github.com/rafa-acioly/urlshor
WORKDIR /go/src/github.com/rafa-acioly/urlshor
RUN go get github.com/lib/pq \
  && go get github.com/go-redis/redis \
  && go get github.com/gorilla/mux \
  && CGO_ENABLED=0 go build -ldflags '-s -w'

FROM scratch
WORKDIR /
COPY --from=builder /go/src/github.com/rafa-acioly/urlshor/urlshor /
ADD ./static /go/src/github.com/rafa-acioly/urlshor/static

CMD ["/urlshor"]
