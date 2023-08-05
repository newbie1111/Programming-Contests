#!/bin/bash

SERVICES_ROOT=("atcoder.jp")

for service in "${SERVICES_ROOT[@]}"; do

    for dirpath in $(find ${service} -type d -regex ".*/test"); do
        if test -d "${dirpath}"; then
            echo "delete directory : $dirpath"
            rm -rf "${dirpath}"
        fi
    done

done
