build:
	docker build -t lomik/go-carbon-build-ubuntu ./

run:
	docker run -ti --rm -v `pwd`:/deb lomik/go-carbon-build-ubuntu

bash:
	docker run -ti --rm -v `pwd`:/deb --entrypoint="/bin/bash" lomik/go-carbon-build-ubuntu
