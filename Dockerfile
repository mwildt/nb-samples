FROM traefik:v2.6

MAINTAINER mail@maltewildt.de

COPY traefik.yaml /etc/traefik/traefik.yaml
COPY file-provider.yaml /etc/traefik/file-provider.yaml