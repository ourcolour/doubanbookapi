FROM centos:7

MAINTAINER CC.Yao <yaochen@xjh.com>

RUN mkdir -p /opt/doubanbookapi
ADD ./bin/doubanbookapi_linux /opt/doubanbookapi
ADD ./resources/favicon.ico /opt/doubanbookapi

CMD ["/opt/doubanbookapi/doubanbookapi_linux"]

EXPOSE 8080