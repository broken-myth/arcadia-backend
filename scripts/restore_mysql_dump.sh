#!/bin/bash

mysqldumpfile=$1
# Pass the appropriate mysqldumpfile as a command line argument

if [ -f .env ]; then
    export $(cat .env | xargs)
fi

docker exec -i arcadia_db mysql -uroot -p${MYSQL_ROOT_PASSWORD} arcadia_23 <${mysqldumpfile}
