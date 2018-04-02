FROM mysql:5.7
ENV MYSQL_ALLOW_EMPTY_PASSWORD=yes
ENV MYSQL_ROOT_PASSWORD=""
COPY ./docker-entrypoint-initdb.d/01_initial_db.sql /docker-entrypoint-initdb.d/01_initial_db.sql
COPY my.cnf /etc/mysql/conf.d/my.cnf
