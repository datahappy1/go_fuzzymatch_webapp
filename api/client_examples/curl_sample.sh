#!/bin/bash

curl --location --request POST 'http://localhost:8080/api/v1/requests/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "stringsToMatch": "'\''apple, gmbh'\'','\''corp'\''",
    "stringsToMatchIn": "hair, bear",
    "mode": "combined"
}'

curl --location --request GET 'http://localhost:8080/api/v1/requests/4027cf00-6ff3-4239-9ec5-2d820b4f93cd/' \
--data-raw ''
