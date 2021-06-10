FROM golang:1.16.5-buster
RUN go get github.com/labstack/echo/v4
ENV GO111MODULE=on
ENV APP_USER app
ENV APP_HOME /go/src/github.com/huzairuje/chatat_backend_engineer
ARG GROUP_ID
ARG USER_ID
RUN groupadd --gid $GROUP_ID app && useradd -m -l --uid $USER_ID --gid $GROUP_ID $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
USER $APP_USER
WORKDIR $APP_HOME
RUN wget -L https://github.com/golang-migrate/migrate/releases/download/v4.8.0/migrate.linux-amd64.tar.gz
RUN tar xvzf migrate.linux-amd64.tar.gz
RUN cp ./migrate.linux-amd64 /bin/migrate
RUN rm migrate.linux-amd64.tar.gz
EXPOSE 3010
CMD ["go", "run"]