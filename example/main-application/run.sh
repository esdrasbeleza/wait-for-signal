#!/bin/sh

echo "Waiting for signal... "
/tmp/wait-for-signal dependency-is-ready 12345

./hello-world
