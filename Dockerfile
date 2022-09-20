FROM python:3-alpine
MAINTAINER sharpezhang@tencent.com
RUN apk update \
    && apk add gcc openssl-dev libffi-dev musl-dev python3-dev \
    && pip install mycli \
    && apk del gcc openssl-dev libffi-dev musl-dev python3-dev \
    && rm -rf /var/cache/apk
ENTRYPOINT [ "mycli" ]