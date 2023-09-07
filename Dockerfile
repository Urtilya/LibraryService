FROM mysql:8

COPY ./library.sql /docker-entrypoint-initdb.d
