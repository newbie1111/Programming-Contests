#!/bin/bash

for arg; do
    oj-prepare ${arg}
    sleep 5
done

return 0
