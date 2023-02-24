FROM alpine

RUN apk add ca-certificates jq

COPY ./build/accountsd /usr/bin/accountsd
COPY ./scripts /scripts

RUN echo "hid"

CMD ["sh", "./scripts/node.sh"]