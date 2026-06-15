cd $(dirname $0)/../docker_test && \
./build.sh && \
cd .. && \
go generate -v ./... && \
go test -v ./... && \
pwd && \
cd integration_test && \
./run_tests.sh
