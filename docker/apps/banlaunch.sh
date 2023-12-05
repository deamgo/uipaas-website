#!/bin/bash
# exec go build -o ./docker/apps/backend ./backend/main.go

exec ./backend -dbConfig ./config.yaml & 

nginx -g "daemon off;"
