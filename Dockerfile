FROM scratch

MAINTAINER mail@maltewildt.de

COPY main /main
COPY static /static

EXPOSE 80
CMD ["/main"]