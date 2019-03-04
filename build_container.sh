#!/bin/sh -e
main=${PWD##*/}
docker pull golang:stretch
docker run --rm \
    -e GOCACHE=off -e GOOS=linux -e GOARCH=amd64 \
    -u `id -u` -v "`pwd`:/go/src/$main" \
    golang:stretch \
    /bin/sh -c "cd /go/src/$main && ./update_dependencies.sh && ./run_unit_tests.sh && go build -a -x -race"
if [ ! -f Dockerfile ]
then
    cat <<EOF > Dockerfile
FROM debian:stretch

RUN apt-get update \\
    && DEBIAN_FRONTEND=noninteractive apt-get upgrade -y \\
    && apt-get install -y \\
        ca-certificates \\
    && apt-get autoremove -y \\
    && apt-get clean \\
    && rm -rf /var/lib/apt/lists/*

ADD $main /opt/

RUN chmod a+w /opt

WORKDIR /opt/
ENTRYPOINT [ "/opt/$main" ]
EOF
fi

docker pull debian:stretch
docker build -t $main .
