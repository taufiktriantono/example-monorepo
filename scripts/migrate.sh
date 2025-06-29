#!/usr/bin/env bash
set -e

DOMAIN=$1      # approval / transaction / user …

DBURL=$2

echo "▶︎ Running migrations for $DOMAIN"

migrate -path "./migrations/$DOMAIN" -database "$DBURL" up
