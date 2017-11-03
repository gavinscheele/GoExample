FROM golang:1.9
ENV GOPATH /go/
RUN mkdir -p /go/src/github.com/pubgslotsExample/src/
ADD src/ /go/src/github.com/pubgslotsExample/src/
WORKDIR /go/src/github.com/pubgslotsExample/src/
ADD Gopkg.lock Gopkg.toml ./
RUN ./build.sh