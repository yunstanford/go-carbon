FROM ubuntu

RUN DEBIAN_FRONTEND=noninteractive apt-get update && apt-get install git make build-essential debhelper vim -y

ADD https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz /var/tmp/
ADD remove-golang-build-depends.patch /opt/remove-golang-build-depends.patch
RUN tar -C /usr/local -xzf /var/tmp/go1.4.2.linux-amd64.tar.gz
RUN ln -s /usr/local/go/bin/go /usr/local/bin/go

ADD build-deb.sh /opt/

ENTRYPOINT ["/opt/build-deb.sh"]

