FROM golang:1.15.8-alpine

# ENV GO111MODULE=on

WORKDIR /go/src/github.com/H-J-Ainashi/CATech-RESTwithGolang

  COPY go.mod go.sum ./
  RUN go mod download

COPY . .
#RUN go mod init github.com/H-J-Ainashi/CATech-RESTwithGolang
#RUN go mod download

# RUN apk add git \
#    && go get -u -t -v github.com/jinzhu/gorm \
#    && go get -u -t -v github.com/labstack/echo \
#    && go get -u -t -v github.com/labstack/echo/middleware

#RUN apk add --no-cache git \
#    && go get -t -v github.com/H-J-Ainashi/CATech-RestwithGolang/domain \
#    && go get -t -v github.com/H-J-Ainashi/CATech-RestwithGolang/UseCase \
#    && go get -t -v github.com/H-J-Ainashi/CATech-RestwithGolang/interfaces/database \
#    && go get -t -v github.com/H-J-Ainashi/CATech-RestwithGolang/interfaces/controllers \
#    && go get -t -v github.com/H-J-Ainashi/CATech-RestwithGolang/infrastructure

RUN go build .

RUN go get github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]
