FROM ubuntu:20.04@sha256:bac6081102aae54ba4bcc714695b8f637e42768c7f376f374c428bab043ddc0f

# FROM nginx:stable
COPY ./server.txt /home/workbench/server.txt

RUN apt-get update --fix-missing && apt-get install -y $(cat /home/workbench/server.txt)

# RUN apt-get update --fix-missing && apt-get install -y nginx && apt-get -y install vim && apt-get install -y golang && apt-get install -y redis-server

ARG APP_HOME

VOLUME ${APP_HOME}
WORKDIR ${APP_HOME}

