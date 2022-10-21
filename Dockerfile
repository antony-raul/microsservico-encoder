FROM alfg/bento4 AS builder

FROM golang:1.17.8-alpine3.15

ENV PATH=/opt/bento4/bin:${PATH}

RUN apk update \
    && apk add --no-cache \
    python3 libgcc gcc g++

COPY --from=builder /opt/bento4 /opt/bento4

WORKDIR /go/src

# vamos mudar para o endpoint correto. Usando top apenas para segurar o processo rodando
ENTRYPOINT [ "top" ]