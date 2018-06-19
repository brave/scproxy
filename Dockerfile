FROM golang
ENV GOPATH /go
RUN curl https://glide.sh/get | sh

WORKDIR /go/src/github.com/brave/scproxy

COPY main.go .
COPY scproxy scproxy
COPY vendor vendor
RUN go install .

CMD /go/bin/scproxy
