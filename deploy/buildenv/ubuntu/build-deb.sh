#!/bin/sh

cd /opt/
git clone https://github.com/lomik/go-carbon.git
cd go-carbon
git tag | tail -n 1 | xargs -I % git checkout % -b %
make submodules
make
patch -p1 < /opt/remove-golang-build-depends.patch
make deb

cp /opt/*.deb /deb/
