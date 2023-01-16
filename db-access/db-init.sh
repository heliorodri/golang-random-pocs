#!/bin/sh

verifyDocker() {
    if [ "$(docker ps -q -f name=$dockername)" ]; then
        if [ "$(docker ps -aq -f status=exited -f name=$dockername)" ]; then
            docker rmi -f $dockername
        else
            docker stop $dockername
            docker container prune $dockername
        fi
    fi
}

verifyDocker "mydb"

docker run --name mydb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql

docker container exec mydb mysql -u root -p root -e ./init-sql.sql

