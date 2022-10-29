#FROM debian:bullseye-slim

FROM ubuntu:kinetic

WORKDIR /app

#ENV PATH="${PATH}:/usr/local/go/bin"

ENV DEBIAN_FRONTEND=noninteractive

RUN apt update -y

RUN apt install -y ca-certificates libtesseract-dev tesseract-ocr build-essential golang
#RUN apt install -y ca-certificates build-essential

#ADD go1.19.linux-amd64.tar.gz  /usr/local
#COPY go.* ./
COPY . .

#COPY mupdf/ ./mupdf

#WORKDIR /app/mupdf

#RUN make HAVE_X11=no HAVE_GLUT=no prefix=/usr/local install

#WORKDIR /app

RUN go mod download

RUN go build .

RUN useradd -m nobody

USER nobody

CMD ["./memo"]
