#!/bin/sh
go build -o send.out send.go 
chmod +x send.out
for filePath in $(find "../" -name "*.c")
do
    ./send.out $filePath $1 &
    sleep 3
done

wait

