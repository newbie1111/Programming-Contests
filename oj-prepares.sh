#!/bin/bash

for arg; do
    echo oj-prepare ${arg} 
done

sleep 5

for arg; do
    oj-prepare ${arg} &
done