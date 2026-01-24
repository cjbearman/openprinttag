FROM ubuntu:24.04

WORKDIR /
COPY source_repo/ /source_repo/
RUN apt update
RUN apt install -y python3-full python3-pip
RUN python3 -m venv /.venv
ENV PATH="/.venv/bin:$PATH:."
WORKDIR /source_repo
RUN pip install -r requirements.txt
WORKDIR /source_repo/utils
RUN chmod -R 755 /source_repo
