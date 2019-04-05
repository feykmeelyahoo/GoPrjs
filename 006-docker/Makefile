
GOVERSION=1.12.1
VERSION=0.0.1

build:
	# GOOS=linux  GOARCH=amd64 go build -o app
    docker run --rm -v "$PWD":/usr/src/kubehostname -w /usr/src/kubehostname golang:"$GOVERSION" go build -v
	docker build -t feykmeel/kubehostname:"$VERSION" .
	rm -f app

push:
	docker push feykmeel/kubehostname:"$VERSION"