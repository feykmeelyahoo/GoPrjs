GOVERSION=1.12.1
VERSION=0.0.2
MYIP := $(shell (ifconfig eth1 | awk '{ print $2}' | grep -E -o "([0-9]{1,3}[\.]){3}[0-9]{1,3}");)
LOCALIP=$(shell ifconfig | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1' | awk '{print $1}'); \

.compile:
	docker run --rm -v ${PWD}:/usr/src/kubehostname -w /usr/src/kubehostname golang:${GOVERSION} go build -v
build: .compile
	# GOOS=linux  GOARCH=amd64 go build -o app
    docker run --rm -v "$PWD":/usr/src/kubehostname -w /usr/src/kubehostname golang:"$GOVERSION" go build -v
	# docker build -t feykmeel/kubehostname:"$VERSION" .
	docker build -t feykmeel/kubehostname:${VERSION} .
	rm -f kubehostname

push:
	docker push feykmeel/kubehostname:${VERSION}
run:
	echo ${MYIP}
	echo ${LOCALIP}
#	for i in `seq 0 1 100`; do curl http://${MYIP}:31540; echo; done
