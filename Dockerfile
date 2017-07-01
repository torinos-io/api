FROM alpine:3.4

# c.f. https://github.com/sgerrand/alpine-pkg-glibc/blob/2.23-r3/README.md
RUN apk add --no-cache \
    ca-certificates gcc wget \
  && wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://raw.githubusercontent.com/sgerrand/alpine-pkg-glibc/master/sgerrand.rsa.pub \
  && wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.23-r3/glibc-2.23-r3.apk \
  && apk add glibc-2.23-r3.apk

RUN apk add --no-cache --update --virtual=deps1 \
    bash \
    coreutils \
    gcc

RUN mkdir /app
COPY ./bin /app/bin
COPY ./script /app/script
COPY ./data /app/data
