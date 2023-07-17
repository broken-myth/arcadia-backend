#!/bin/sh

until nc -z -v -w30 arcadia_23_db 3306; do
   echo "Waiting for database connection..."
   sleep 5
done

chmod 777 -R logs/ docker_volumes/ # Permission issues

sleep 3

echo -e "\e[34m >>> Starting the server \e[97m"
$1
