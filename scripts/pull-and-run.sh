#!/bin/bash

LOCAL=$(git log --oneline | head -1)
REMOTE=$(git log --oneline origin/main | head -1)

cd ~/hauslab

if [[ $LOCAL != $REMOTE ]]; then
    git pull
    cd ../services && make up
fi
