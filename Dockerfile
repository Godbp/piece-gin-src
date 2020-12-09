# pengbo golang-gin
# Version 1.0.0

# Base Image

FROM ubuntu

MAINTAINER Godbp

ENV PATH /root/mysql/config

RUN mkdir -p /root/mysql/config
RUN cd /root/mysql/config
RUN touch mysql.config
RUN echo -e "" >> /root/mysql/config/mysql.config