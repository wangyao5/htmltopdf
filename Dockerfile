FROM ubuntu

MAINTAINER 82486240@qq.com
ENV TIMEZONE="Asia/Shanghai"
ENV DEBIAN_FRONTEND="noninteractive"
WORKDIR /opt/

RUN apt-get update --fix-missing
RUN apt-get install -y tzdata && \
ln -fs /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
dpkg-reconfigure -f noninteractive tzdata

# Install dependencies needed for wkhtmltopdf
RUN apt-get install -y libxrender1 libfontconfig1 libxext6 fonts-arphic-bkai00mp fonts-arphic-bsmi00lp \
 fonts-arphic-gbsn00lp fonts-arphic-gkai00mp fonts-dejavu-core fonts-droid-fallback fonts-liberation fonts-lmodern \
 fonts-tibetan-machine xfonts-utils xfonts-75dpi xfonts-100dpi

RUN apt-get install -y wkhtmltopdf

COPY main .
CMD ["/opt/main"]

