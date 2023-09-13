#!/bin/bash

for arg; do
    echo oj-prepare ${arg} 
done

sleep 3

for arg; do
    oj-prepare ${arg}
    sleep 1
done