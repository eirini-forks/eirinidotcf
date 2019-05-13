#!/bin/bash

cat <<EOF
{
  "db_name": "pheed",
  "username":"root",
  "password":"$MYSQL_ROOT_PASSWORD",
  "db_address": "$DB_IP:$DB_PORT"
}
EOF
