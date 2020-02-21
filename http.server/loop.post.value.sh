#!/bin/bash

port=$1;
value=$2;

if [ "$port" = "" ]; then echo;echo "need port";echo;echo; exit 1; fi;
if [ "$value" = "" ]; then echo;echo "need value";echo;echo;  exit 1; fi;

while [ 1 ];
do

    curl -X POST -s "http://localhost:${port}/value/{$value}"

done;
