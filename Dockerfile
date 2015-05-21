FROM golang
MAINTAINER Brian Morton "brian@mmm.hm"

RUN apt-get -y update && \
    apt-get -y install libpcap0.8-dev

ENV REPO github.com/bmorton/dnswatch
ADD . /go/src/$REPO
RUN cd /go/src/$REPO && go get -v -d
RUN go install $REPO

CMD ["/go/bin/dnswatch"]
