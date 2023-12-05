FROM ubuntu:latest

FROM nginx:stable

RUN apt-get update && apt-get -y install vim && apt-get install -y golang

ARG APP_HOME

VOLUME ${APP_HOME}
WORKDIR ${APP_HOME}

