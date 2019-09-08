FROM mysql:5.7
 
RUN apt-get update && \
    apt-get install -y apt-utils locales vim && \
    rm -rf /var/lib/apt/lists/* && \
    echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen && \
    locale-gen ja_JP.UTF-8
ENV LC_ALL ja_JP.UTF-8

ADD ./docker/db/charset.cnf /etc/mysql/conf.d/charset.cnf
ADD ./docker/db/log.cnf /etc/mysql/conf.d/log.cnf
ADD ./docker/db/init.sql /docker-entrypoint-initdb.d/init.sql
