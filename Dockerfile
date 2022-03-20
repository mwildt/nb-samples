FROM certbot/certbot

MAINTAINER mail@maltewildt.de

WORKDIR /

COPY start.sh /start.sh
COPY main /main
COPY static /static

RUN mkdir -p /www/.well-known/

ENTRYPOINT [ "/bin/sh", "-l", "-c" ]
CMD ["/main"]