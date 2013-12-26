FROM centos

MAINTAINER mimoto-xxxxxx

RUN \
 rpm --import \
  http://dl.fedoraproject.org/pub/epel/RPM-GPG-KEY-EPEL-6 \
  http://apt.sw.be/RPM-GPG-KEY.dag.txt; \
 rpm -ivh \
  http://dl.fedoraproject.org/pub/epel/6/x86_64/epel-release-6-8.noarch.rpm \
  http://apt.sw.be/redhat/el6/en/x86_64/rpmforge/RPMS/rpmforge-release-0.5.2-2.el6.rf.x86_64.rpm; \

 yum install -y --enablerepo=rpmforge-extras \
  git \
  hg

RUN \
 cd /root; \
 wget -nv https://go.googlecode.com/files/go1.2.linux-amd64.tar.gz; \
 tar zxf go1.2.linux-amd64.tar.gz; \
 cd go; \
 mkdir gopath; \
 export GOROOT=/root/go; \
 export GOPATH=$GOROOT/gopath; \
 export PATH=$PATH:$GOROOT/bin:$GOPATH/bin; \
 go get bitbucket.org/tebeka/selenium; \
 rm /root/go1.2.linux-amd64.tar.gz

ADD ./files /

CMD bash /root/startup.bash

