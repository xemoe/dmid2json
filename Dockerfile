FROM golang:latest

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN git config --global http.sslVerify false
RUN git config --global http.postBuffer 1048576000
RUN apt-get update -yq && apt-get install -yq dmidecode
