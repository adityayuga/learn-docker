#!/bin/sh

postgres_ready() {
  $(which curl) http://$POSTGRES_HOST:$POSTGRES_PORT/ 2>&1 | grep '52'
}

until postgres_ready; do
  >&2 echo 'Waiting for PostgreSQL to become available...'
  sleep 1
done
>&2 echo 'PostgreSQL is available.'

cmd="$@"
exec $cmd