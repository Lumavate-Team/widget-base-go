from golang:alpine

RUN apk add --no-cache \
    git \
    curl \
  && mkdir -p /go/src/widget

WORKDIR /go/src/widget
COPY ./widget /go/src/widget

RUN go get -u github.com/astaxie/beego
RUN go get -u github.com/beego/bee
RUN go get -u github.com/Lumavate-Team/go-signer
RUN go get -u github.com/Lumavate-Team/go-properties

CMD [ "bee", "run" ]
