#!/bin/bash

if [ -z "${USERNAME}" ]; then
    echo -e "env 'USERNAME' is required."
    exit 1
fi
if [ -z "${PASSWORD}" ]; then
    echo -e "env 'PASSWORD' is required."
    exit 1
fi
if [ -z "${HOST}" ]; then
    echo -e "env 'HOST' is required."
    exit 1
fi
if [ -z "${PORT}" ]; then
    echo -e "env 'PORT' is required."
    exit 1
fi
if [ -z "${DBNAME}" ]; then
    echo -e "env 'PORT' is required."
    exit 1
fi

echo -e ""

sam package \
    --template-file template.yaml \
    --s3-bucket my-blog-api-sam \
    --output-template-file packaged.yaml

sam deploy \
    --template-file packaged.yaml \
    --stack-name my-blog-api-sam \
    --capabilities CAPABILITY_IAM \
    --parameter-overrides \
        USERNAME=${USERNAME} \
        PASSWORD=${PASSWORD} \
        HOST=${HOST} \
        PORT=${PORT} \
        DBNAME=${DBNAME}