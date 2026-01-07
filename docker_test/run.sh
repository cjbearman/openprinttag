#!/bin/bash

CMDFILE=./shared/.command

trap "rm $CMDFILE" EXIT
echo "cd /source_repo/utils" > $CMDFILE
echo $* >> $CMDFILE
chmod 755 $CMDFILE

docker run -it --rm --user $(id -u):$(id -g) --mount type=bind,src=./shared,dst=/shared opttest:latest bash /shared/.command
