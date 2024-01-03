FROM ubuntu:20.04@sha256:080169816683e6f063d3903434565624287828ecfd06bd2f813b30325e8b1eca

# FROM nginx:stable
COPY ./server.txt /home/workbench/server.txt

RUN apt-get update --fix-missing && apt-get install -y $(cat /home/workbench/server.txt)

# RUN apt-get update --fix-missing && apt-get install -y nginx && apt-get -y install vim && apt-get install -y golang && apt-get install -y redis-server

ARG APP_HOME

VOLUME ${APP_HOME}
WORKDIR ${APP_HOME}

