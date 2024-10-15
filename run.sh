#!/bin/bash

BINARY_FILE_PATH="bin/main"
MAIN_FILE_PATH="todo-app/main/main.go"
ENV_FILE="config.env"

set -e
ROOT_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

if [ -f "${ROOT_PATH}/${ENV_FILE}" ]; then
    export $(grep -v '^#' ${ROOT_PATH}/${ENV_FILE} | xargs)
fi

while [[ $# -gt 0 ]]
do
    key="$1"
       case ${key} in
            --debug)
                DEBUG=true
                shift
            ;;
            --*)
                echo "Unknown flag $1"
                exit 1
            ;;
    esac
done

function cleanup() {
    if [[ ${DEBUG} == true ]]; then
       echo "Cleaning up..."
       rm  ${ROOT_PATH}/${BINARY_FILE_PATH}
    fi
}

trap cleanup EXIT

if [[ -z ${DEBUG_PORT} ]]; then
    DEBUG_PORT=40000
fi

if [[ ${DEBUG} == true ]]; then
    echo "Running in debug mode on port ${DEBUG_PORT}..."
    CGO_ENABLED=0 go build -gcflags="all=-N -l" -o ${ROOT_PATH}/${BINARY_FILE_PATH} ${ROOT_PATH}/${MAIN_FILE_PATH}
    dlv --listen=:$DEBUG_PORT --headless=true --api-version=2 exec ${ROOT_PATH}/${BINARY_FILE_PATH}
else
    echo "Running in normal mode..."
    go run ${ROOT_PATH}/${MAIN_FILE_PATH}
fi