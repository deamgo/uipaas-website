#!/bin/bash

max_attempts=10
attempts=0

redis-server &

while [ $attempts -lt $max_attempts ]; do
    ./backend -dbConfig ./config.yaml &
    sleep 10
    if pgrep -x "backend" > /dev/null; then
        echo "Backend is running."
        break
    else
        echo "Backend not running. Retrying..."
        sleep 5
        ((attempts++))
    fi
done

if [ $attempts -eq $max_attempts ]; then
    echo "Error: Could not start backend after $max_attempts attempts."
    exit 1
fi

nginx -g "daemon off;"
