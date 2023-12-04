FROM centos:7

FROM nginx:stable

RUN apt-get update && apt-get -y install vim

ARG APP_HOME

# RUN mkdir "/home/uipaashome/backend"

# VOLUME /home/uipaashome/backend
VOLUME ${APP_HOME}
# WORKDIR /home/uipaashome/backend
WORKDIR ${APP_HOME}

ENTRYPOINT [ "./backend","-dbConfig","./config.yaml" ]

