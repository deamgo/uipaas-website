FROM centos@SHA256:a1801b843b1b

# FROM nginx:stable
COPY ./server.txt /home/workbench/server.txt

RUN yum -q clean expire-cache && yum -q update --setopt=obsoletes=1 && yum -q install -y $(cat /home/workbench/server.txt)

# RUN apt-get update --fix-missing && apt-get install -y nginx && apt-get -y install vim && apt-get install -y golang && apt-get install -y redis-server

ARG APP_HOME

VOLUME ${APP_HOME}
WORKDIR ${APP_HOME}

