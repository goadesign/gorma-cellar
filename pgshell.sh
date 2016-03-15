#! /bin/bash

set -e

docker run -it --link gorma-cellar-db:postgres --rm postgres sh -c 'exec psql -h "$POSTGRES_PORT_5432_TCP_ADDR" -p "$POSTGRES_PORT_5432_TCP_PORT" -U gorma'

