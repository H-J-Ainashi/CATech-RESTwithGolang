FROM golang:1.15.8

ENV GO111MODULE=on

WORKDIR /go/src/github.com/H-J-Ainashi/CATech-RESTwithGolang

# COPY go.mod go.sum ./

COPY . .
RUN go mod init github.com/H-J-Ainashi/CATech-RESTwithGolang
RUN go mod download

RUN go build .

RUN go get github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]
