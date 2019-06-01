FROM golang:1.12

COPY ./api-db-migrate.sh /usr/local/bin/api-db-migrate
RUN chmod +x /usr/local/bin/api-db-migrate

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v3.5.4/migrate.linux-amd64.tar.gz | tar xvz
RUN mv -v migrate.linux-amd64 /usr/local/bin/migrate
