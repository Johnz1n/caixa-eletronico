
FROM golang:1.14-alpine

COPY . $GOPATH/src/caixa-eletronico

ENV GIN_MODE=release

RUN apk add --no-cache bash git openssh gcc libc-dev

WORKDIR $GOPATH/src/caixa-eletronico
RUN go get
RUN go install

EXPOSE 8080

CMD [ "caixa-eletronico" ]