FROM golang
WORKDIR /go/src/github.com/RyanJarv/scproxy
RUN curl https://glide.sh/get | sh
CMD go run main.go