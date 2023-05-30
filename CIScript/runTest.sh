#!/bin/sh
for filePath in $(find "../" -name "*.c")
do
    ./send.out $filePath $1 &
done

wait

