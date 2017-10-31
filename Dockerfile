FROM golang:1.9
ENV GOPATH /go/
RUN mkdir -p /go/src/github.com/gdscheele/Example/src/
ADD src/ /go/src/github.com/gdscheele/Example/src/
WORKDIR /go/src/github.com/gdscheele/Example/src/
ADD Gopkg.lock Gopkg.toml ./
RUN ./build.sh