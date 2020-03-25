#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
   if [ "$( psql -tAc "SELECT 1 FROM pg_database WHERE datname='konga'" )" = '1' ]
then
    echo "Database already exists"
else
    creatdb konga
fi
EOSQL