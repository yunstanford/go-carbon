build:
	docker build -t lomik/go-carbon-build-centos6 ./

run:
	docker run -ti --rm -v `pwd`:/rpm lomik/go-carbon-build-centos6

bash:
	docker run -ti --rm -v `pwd`:/rpm --entrypoint="/bin/bash" lomik/go-carbon-build-centos6
