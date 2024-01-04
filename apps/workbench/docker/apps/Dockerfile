FROM amd64/ubuntu:20.04

COPY ./server.txt /home/workbench/server.txt

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update --fix-missing && apt-get install -y $(cat /home/workbench/server.txt)

ARG APP_HOME

VOLUME ${APP_HOME}
WORKDIR ${APP_HOME}