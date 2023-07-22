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

        if test $SUBMIT_FILE_EXT = "py"; then
            # oj s --guess-python-interpreter pypy $SUBMIT_FILE -y
            oj s $SUBMIT_FILE -y
        else
            oj s $SUBMIT_FILE -y
        fi

        sleep 5
    else
        echo $SUBMIT_FILE is not exist.
    fi

    cd $SCRIPT_DIR

done

return 0
