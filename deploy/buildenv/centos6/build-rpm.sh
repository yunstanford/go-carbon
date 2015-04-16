#!/bin/sh

cd /opt/
git clone https://github.com/lomik/go-carbon.git
cd go-carbon
git tag | tail -n 1 | xargs -I % git checkout % -b %
make submodules
make
make rpm

cp tmp/rpm/RPMS/x86_64/*.rpm /rpm/
