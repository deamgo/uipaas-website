FROM centos:7

ARG APP_HOME

# RUN mkdir "/home/uipaashome/backend"

# VOLUME /home/uipaashome/backend
VOLUME ${APP_HOME}
# WORKDIR /home/uipaashome/backend
WORKDIR ${APP_HOME}

ENTRYPOINT [ "./backend","-dbConfig","./config.yaml" ]