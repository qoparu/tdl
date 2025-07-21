#!/bin/sh
/wait-for-it.sh db:5432
/wait-for-it.sh mqtt:1883
exec /server