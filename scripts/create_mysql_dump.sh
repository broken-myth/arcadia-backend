#!/bin/bash

if [[ $(docker ps --filter status=running | grep arcadia_db) ]]
then
    mkdir -p mysql_dumps
    time=$(date +%d-%m-%Y-%H-%M-%S)

    if [ -f .env ]
    then
        export $(cat .env | xargs)
    fi
    docker exec arcadia_db /usr/bin/mysqldump -u root -p${MYSQL_ROOT_PASSWORD} arcadia_23 > ./mysql_dumps/dump-$time.sql
    
fi
