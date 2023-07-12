#!/bin/bash

mysqldumpfile="mysql_dumps/dump-21-12-2022-19-13-34.sql" 
# Change this to the appropriate mysqldumpfile

if [ -f .env ]
then
    export $(cat .env | xargs)
fi

docker exec -i arcadia_db mysql -uroot -p${MYSQL_ROOT_PASSWORD} arcadia_23 < ${mysqldumpfile};
