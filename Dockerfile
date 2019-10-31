FROM alpine:latest
MAINTAINER mohammad hewedy

ENV APP_LOCATION /app/Cloud-Native-Go

COPY Cloud-Native-Go ${APP_LOCATION}
RUN chmod +x ${APP_LOCATION}

ENV PORT 8080
EXPOSE ${PORT}

ENTRYPOINT ${APP_LOCATION}
