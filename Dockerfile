FROM alpine:latest

MAINTAINER Ryan Sloan <rysloan4@gmail.com>

WORKDIR "/opt"

ADD .docker_build/chat /opt/bin/chat
ADD ./templates /opt/templates

CMD ["/opt/bin/go-getting-started"]
