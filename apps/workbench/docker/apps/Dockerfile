FROM arm64v8/ubuntu:22.04

# FROM nginx:stable
COPY ./server.txt /home/workbench/server.txt

RUN apt-get update --fix-missing && apt-get install -y $(cat /home/workbench/server.txt) && apt-get -qq update \
    && apt-get -qq install -y --no-install-recommends ca-certificates curl


# RUN apt-get update --fix-missing && apt-get install -y nginx && apt-get -y install vim && apt-get install -y golang

ARG APP_HOME

VOLUME ${APP_HOME}
WORKDIR ${APP_HOME}

