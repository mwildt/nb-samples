FROM ubuntu:latest

MAINTAINER mail@maltewildt.de

COPY main /main
COPY static /static

CMD ["/main"]