#!/bin/sh

echo "Running some slow task..."

for i in `seq 10 -1 1`; do 
    echo "$i..."
    sleep 1
done

echo "Done!"

/tmp/send-signal dependency-is-ready main-application:12345
