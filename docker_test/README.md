# Docker test environment for prusa python tools

Allows for creating a test container with everything needed to run prusa python based open print tag tools in a repeatable, consistent, isolated fashion without having to install local dependencies.

## Build the container

Run build.sh FROM THIS DIRECTORY to create a docker container with the prusa python distro and all dependencies installed

## Use the container

Using run.sh FROM THIS DIRECTORY you can invoke the python tools.
For example:

```
./run.sh 'python nfc_initialize.py -s 304 | python rec_update.py /shared/data_to_fill.yaml | python rec_info.py --show-all --opt-check'
```	
