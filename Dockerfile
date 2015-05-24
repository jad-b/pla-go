FROM ubuntu:14.04

RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    curl \
    gcc \
    git \
    make

# Install Go
ENV GO_VERSION 1.4.2
ENV OS linux
ENV ARCH amd64
# Download and extract the binary to /usr/local/go/{src,bin,pkg}
RUN curl -sSL https://storage.googleapis.com/golang/go$GO_VERSION.$OS-$ARCH.tar.gz \
    | tar -v -C /usr/local/ -xz && \
    mkdir -p /go/bin
ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

# Mount you $GOPATH/src directory
VOLUME /go/src
WORKDIR /go/src
