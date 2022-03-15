FROM python:latest

MAINTAINER mail@maltewildt.de

COPY index.html /

EXPOSE 7000
CMD python -m http.server 7000