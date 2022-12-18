FROM golang:1.19

RUN mkdir -p /usr/src/sijl
COPY . /usr/src/sijl
WORKDIR /usr/src/sijl
RUN go mod tidy

RUN apt-get update
RUN apt-get -y install python3
RUN apt-get -y install python3-setuptools
RUN apt-get -y install python3-pip
RUN pip install Pillow
RUN pip install numpy

ENTRYPOINT go run cmd/sijl.go
