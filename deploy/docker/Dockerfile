ARG GO_VERSION=1.20.2
FROM ubuntu:20.04

COPY . /go/webservice

RUN /bin/bash -c "cd /go/webservice/ \
     && apt-get update \
     && apt-get install -y make wget git \
     && wget https://golang.org/dl/go1.21.5.linux-amd64.tar.gz \
     && tar -zxvf go1.21.5.linux-amd64.tar.gz -C /usr/local \
     && rm -rf go1.21.5.linux-amd64.tar.gz \
     && ln -s /usr/local/go/bin/go /usr/bin/go \
     && make \
     && cp -r ./_build/webservice-* /opt/webservice \
     && rm -rf /go \
     && rm -rf /root/go \
     && apt-get clean \
     && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*"

WORKDIR /opt/webservice

CMD ["/opt/webservice/bin/webservice", "run"]

EXPOSE 8088