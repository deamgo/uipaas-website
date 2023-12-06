FROM ubuntu:22.04

# FROM nginx:stable

RUN apt-get update --fix-missing && apt-get install -y nginx && apt-get -y install vim && apt-get install -y golang

ARG APP_HOME

VOLUME ${APP_HOME}
WORKDIR ${APP_HOME}

