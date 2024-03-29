#!/bin/sh
# wait-for-postgres.sh

set -e

shift
cmd="$@"

until PGPASSWORD=wbpass psql -h "db" -U "wbuser" -d "wbdb" -c '\q'; do
    >&2 echo "Postgres is unavailable - sleeping"
    sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd