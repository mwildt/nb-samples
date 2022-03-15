FROM scratch

MAINTAINER mail@maltewildt.de

COPY main /main
COPY static /static

EXPOSE 3000
CMD ["/main"]