FROM certbot/certbot

MAINTAINER mail@maltewildt.de

RUN mkdir -p /opt/mwcertbot
WORKDIR /opt/mwcertbot

COPY start.sh /opt/mwcertbot/start.sh
COPY main /opt/mwcertbot/main

RUN mkdir -p /opt/mwcertbot/.well-known/

ENTRYPOINT [ "/opt/mwcertbot/main" ]

CMD []