FROM centos:centos6

RUN yum install -y git tar rpm-build make

ADD https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz /var/tmp/
RUN tar -C /usr/local -xzf /var/tmp/go1.4.2.linux-amd64.tar.gz
RUN ln -s /usr/local/go/bin/go /usr/local/bin/go

ADD build-rpm.sh /opt/

ENTRYPOINT ["/opt/build-rpm.sh"]

