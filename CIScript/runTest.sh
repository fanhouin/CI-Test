#!/bin/sh
for filePath in $(find "../" -name "*.c")
do
    go run send.go $filePath $1 &
    sleep 3
done

sleep 190

