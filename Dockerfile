FROM golang

COPY ./app /go/src/github.com/sangwoo89118/go_practice/app
WORKDIR /go/src/github.com/sangwoo89118/go_practice/app

RUN go get ./
RUN go build

RUN go get github.com/pilu/fresh

CMD fresh

EXPOSE 8000