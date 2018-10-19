FROM centos:7

MAINTAINER CC.Yao <yaochen@xjh.com>

ADD ./bin/doubanbookapi_linux /

CMD ["/doubanbookapi_linux"]

EXPOSE 8080