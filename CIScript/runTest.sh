#!/bin/sh
for filePath in $(find "../" -name "*.c")
do
    ./send.out $filePath $1 &
    sleep 3
done

wait

