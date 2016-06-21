FROM alpine:latest

MAINTAINER Ryan Sloan <rysloan4@gmail.com>

WORKDIR "/opt"

ADD .docker_build/main /opt/bin/main
ADD ./templates /opt/templates

CMD ["/opt/bin/main"]
