#!/bin/bash

BINARY_FILE_PATH="bin/main"
MAIN_FILE_PATH="library-app/main/main.go"
ENV_FILE=".env"

set -e
ROOT_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

if [ -f "${ROOT_PATH}/${ENV_FILE}" ]; then
    export $(grep -v '^#' ${ROOT_PATH}/${ENV_FILE} | xargs)
fi

PRESERVE_DB=false
REUSE_DB=false

while [[ $# -gt 0 ]]
do
    key="$1"
       case ${key} in
            --debug)
                DEBUG=true
                shift
            ;;
            --preserve-db)
                PRESERVE_DB=true
                shift
            ;;
            --*)
                echo "Unknown flag $1"
                exit 1
            ;;
    esac
done

POSTGRES_CONTAINER="library-app-db"
CONNECTION_STRING="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSL}"

function cleanup() {
    if [[ ${DEBUG} == true ]]; then
       echo "Cleaning up..."
       rm  ${ROOT_PATH}/${BINARY_FILE_PATH}
    fi

    if [[ ${PRESERVE_DB} = false ]]; then
        echo "Remove DB container"
        docker rm --force ${POSTGRES_CONTAINER}
    else
        echo "Keeping DB container running"
    fi
}

trap cleanup EXIT

if [[ ${REUSE_DB} = true ]]; then
    echo "DB is reused."
else
    set +e
    echo "Create DB container ${POSTGRES_CONTAINER} ..."
    docker run -d --name ${POSTGRES_CONTAINER} \
                -e POSTGRES_HOST=${POSTGRES_HOST} \
                -e POSTGRES_PORT=${POSTGRES_PORT} \
                -e POSTGRES_DB=${POSTGRES_DB} \
                -e POSTGRES_USER=${POSTGRES_USER} \
                -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
                -p ${POSTGRES_PORT}:${POSTGRES_PORT} \
                postgres:${POSTGRES_VERSION}

#    if [[ $? -ne 0 ]] ; then
#        PRESERVE_DB=true
#        exit 1
#    fi

    echo "Waiting for DB to start ..."
    for i in {1..10}
    do
        docker exec ${POSTGRES_CONTAINER} pg_isready -U "${POSTGRES_USER}" -h "${POSTGRES_HOST}" -p "${POSTGRES_PORT}" -d "${POSTGRES_DB}"
        if [ $? -eq 0 ]
        then
            DB_STARTED=true
            break
        fi
        sleep 1
    done

    if [ "${DB_STARTED}" != true ] ; then
        echo "DB in container ${POSTGRES_CONTAINER} was not started. Exiting."
        exit 1
    fi
    set -e

    echo "Execute DB Migrations ..."
    migrate -path ${ROOT_PATH}/db/migrations -database "${CONNECTION_STRING}" up

#    echo "Load initial content ..."
#    ls -d ${ROOT_PATH}/db/init/* | sort | xargs -I {} cat {} |\
#        docker exec -i ${POSTGRES_CONTAINER} psql -U "${POSTGRES_USER}" -h "${POSTGRES_HOST}" -p "${POSTGRES_PORT}" -d "${POSTGRES_DB}"
fi

echo "Migration version: $(migrate -path ${ROOT_PATH}/db/migrations -database "${CONNECTION_STRING}" version 2>&1)"

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