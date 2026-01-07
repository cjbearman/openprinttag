#!/bin/bash

if [ "${VENV_DIR}" != "" ]; then
	echo "Virtual environment already activated"
else 


	PROJ_DIR=$(dirname $0)/..
	VENV_DIR=${PROJ_DIR}/.venv
	if [ ! -d ${VENV_DIR} ]; then
		# Virtual environment not found, create it
		PYTHON=python
		which python3 >/dev/null 2>/dev/null
		if [ $? -eq 0 ]; then
			PYTHON=python3
		fi
		${PYTHON} -m venv ${VENV_DIR}
	fi

	source ${VENV_DIR}/bin/activate

	REPO_DIR=${PROJ_DIR}/source_repo
	cd ${REPO_DIR}
	pip install -r requirements.txt

fi

