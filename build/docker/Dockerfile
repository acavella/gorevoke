ARG BUILD_TAG=v1.0.0-rc12

FROM docker.io/library/golang:1.20.12-alpine3.19 as builder
ENV GOREVOKE_TAG=${BUILD_TAG}
WORKDIR /go/src/

RUN apk add --no-cache -U git \
    && git clone https://github.com/acavella/gorevoke \
    && cd gorevoke && git checkout ${GOREVOKE_TAG} \
    && /bin/sh /go/src/gorevoke/build/docker/bin/install.sh

FROM registry.access.redhat.com/ubi9/ubi:9.3-1476 as runner
ENV GOREVOKE_TAG=${BUILD_TAG}

LABEL maintainer="Tony Cavella <tony@cavella.com>"
LABEL app.version=${GOREVOKE_TAG}

COPY --from=builder /usr/local/bin/gorevoke/ /usr/local/bin/gorevoke/

RUN dnf update -y

EXPOSE 4000
CMD [ "/usr/local/bin/gorevoke/gorevoke" ]
