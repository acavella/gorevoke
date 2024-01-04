FROM docker.io/library/golang:1.20.12-alpine3.19 as builder

ENV GOREVOKE_TAG=v1.0.0-rc.1
WORKDIR /go/src/github.com/gorevoke/

RUN apk get --no-cache -U libc-dev curl nodejs npm git gcc g++ \
    && git clone https://github.com/acavella/gorevoke \
    && cd gorevoke && git checkout ${GOREVOKE_TAG} \
    && ./deployment/docker/prod/bin/install

FROM registry.access.redhat.com/ubi9/ubi:9.3-1476 as runner
LABEL maintainer="Tony Cavella <tony@cavella.com>"

COPY --from=builder /usr/local/bin/semaphore-wrapper /usr/local/bin/
COPY --from=builder /usr/local/bin/semaphore /usr/local/bin/
COPY ./systemd/ /etc/systemd/system/
COPY ./nginx.conf /etc/nginx/conf.d/
COPY ./ansible/ /app/ansible/

RUN --mount=type=secret,id=gitssh dnf update -y && dnf install -y git python3-devel python3-pip nginx rust nc \
    && pip3 install --upgrade pip \
    && pip3 install setuptools-rust ansible paramiko \
    && systemctl enable nginx semaphore first-run \
    && mkdir -p ~/.ssh \
    && ssh-keyscan -t rsa github.boozallencsn.com > ~/.ssh/known_hosts \
    && echo -e "Host *\nIdentityFile /run/secrets/gitssh" > ~/.ssh/config \
    && ansible-galaxy collection install community.general \
    && /usr/local/bin/ansible-playbook -vv /app/ansible/setup.yml \
    && rm -Rf ~/.ssh

EXPOSE 4000
CMD [ "/opt/gorevoke/gorevoke" ]