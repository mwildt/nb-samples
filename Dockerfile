FROM mongo:latest

MAINTAINER mail@maltewildt.de

CMD ["mongod --smallfiles  --optlogsize 128"]