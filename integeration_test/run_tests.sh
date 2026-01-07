#! /bin/bash

function runtest {
    set -e
    TESTNAME=$1

    echo -n "Running ${TESTNAME}..."

    # First we run the commands in "<testname>.docker.commands" in docker
    # These should produce <testname>.docker.python.bin
    docker run -it --rm --user $(id -u):$(id -g) --mount type=bind,src=.,dst=/shared opttest:latest bash /shared/$TESTNAME.docker.commands

    # Now we run the commands in "<testname>.golang.commands"
    # These should produce <testname>.golang.bin
    bash $TESTNAME.golang.commands

    # Make hexidecimal copies of both files, because it sucks to debug binary
    cat $TESTNAME.python.bin | gohexdump > $TESTNAME.python.hex
    cat $TESTNAME.golang.bin | gohexdump > $TESTNAME.golang.hex
    set +e    

    # Now we can look for differences and report as needed
    if diff $TESTNAME.python.bin $TESTNAME.golang.bin >/dev/null 2>&1; then
        echo "Pass"
    else
        echo "Fail"
        echo "Python out:"
        cat $TESTNAME.python.hex
        echo "Golang out:"
        cat $TESTNAME.golang.hex
        echo "Diff"
        diff $TESTNAME.python.hex $TESTNAME.golang.hex
    fi
}

# Script starts here
cd $(dirname $0)
BASEDIR=$(pwd)

# Check for docker
if ! which docker >/dev/null 2>&1; then
    echo "Can't find docker"
    exit 1
fi

# Check to make sure we have our docker image, and build it not
if [ -z "$(docker image ls --format '{{.ID}}' opttest:latest)" ]; then
    cd ../docker_test
    ./build.sh
fi

# Build our required binaries, ensuring we exit on failure
set -e
cd $BASEDIR/../cmd/optag
go build -o ${BASEDIR}/optag ./optag.go
set +e

cd $BASEDIR

# Now we can run each test
runtest test1
runtest test2
runtest test3
