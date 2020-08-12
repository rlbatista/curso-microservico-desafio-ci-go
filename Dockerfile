FROM golang:rc-alpine3.12
WORKDIR /src/sum_go
COPY ./src/sum_go/. .