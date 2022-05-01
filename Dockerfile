#Go API
FROM golang:latest AS builder
ADD . /app
WORKDIR /app/server
RUN go mod download
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -ldflags "-w" -a -o /main .

#Build React application
FROM node:alpine AS node_builder
COPY --from=builder /app/client ./
RUN npm install
RUN npm run build

#Final Stage, this will be container
FROM ubuntu:18.04

##################################################
# install golang-migrate
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y curl gnupg2 vim && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz | tar xvz && \
    mv ./migrate.linux-amd64 /usr/bin/migrate

# install mysql-client
RUN apt-get update &&\
    apt-get install -y mysql-client

ENV MYSQL_URL='mysql://root:root@tcp(mysql:3306)/react_go_app?multiStatements=true'

COPY ./mysql/000001_init.up.sql /root/migrations/example1/000001_init.up.sql
COPY ./mysql/000001_init.down.sql /root/migrations/example1/000001_init.down.sql
# migrate -database ${MYSQL_URL} -path migrations/example1 up
##################################################

COPY --from=builder /main ./server/
COPY --from=node_builder /build ./client/build
WORKDIR /server
# 使うPORTをENVで設定する これ意味案のか?
#ENV PORT=8080

RUN chmod +x ./main
EXPOSE 8080

CMD ./main