#!/bin/bash

for arg; do
    oj-prepare ${arg} &
done

return 0
