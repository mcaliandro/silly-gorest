FROM golang:1.20-alpine

ENV PORT=8080

RUN mkdir /app
COPY bin/silly-gorest /app/silly-gorest

EXPOSE ${PORT}
ENTRYPOINT [ "/app/silly-gorest" ]