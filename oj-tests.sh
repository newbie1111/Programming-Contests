#!/bin/bash

SCRIPT_DIR=$(
    cd $(dirname $0)
    pwd
)

for arg; do
    SUBMIT_DIR=$(dirname ${arg})
    SUBMIT_FILE=${arg##*/}
    SUBMIT_FILE_EXT=${arg##*.}

    if test -d $SUBMIT_DIR; then
        cd $SUBMIT_DIR
    fi

    if test -f $SUBMIT_FILE; then

        if test $SUBMIT_FILE_EXT = "go"; then
            oj t -c "go run main.go"
        elif test $SUBMIT_FILE_EXT = "py"; then
            oj t -c "python main.py"
        elif test $SUBMIT_FILE_EXT = "cpp"; then
            g++ main.cpp
            oj t
            rm a.out
        else
            echo $SUBMIT_FILE_EXT is not supported.
        fi

    else
        echo $SUBMIT_FILE is not exist.
    fi

    cd $SCRIPT_DIR
done
