from golang:alpine

RUN apk add --no-cache \
                git \
                curl \
#&& git rev-parse HEAD > /revision \
#&& rm -rf .git \
  && mkdir -p /go/src/widget

WORKDIR /go/src/widget
COPY ./widget /go/src/widget

RUN go get -u github.com/astaxie/beego
RUN go get -u github.com/beego/bee
RUN go get -u github.com/Lumavate-Team/go-signer

CMD bee run
