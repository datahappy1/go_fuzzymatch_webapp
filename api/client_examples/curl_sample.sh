#!/bin/bash

# POST REQUEST
# *Nix terminal:
curl --location --request POST 'http://localhost:8080/api/v1/requests/' \
--header 'Content-Type: text/plain' \
--data-raw '{
  "stringsToMatch": "'\''231 Beechwood Street'\'', '\''Helena, MT 59601'\'', '\''866 Brook Court'\'', '\''Harrison Township, MI 48045'\'', '\''40 Bayport Street'\'', '\''Virginia Beach, VA 23451'\'', '\''20 Hanover St.",
	"stringsToMatchIn": "'\''231 Beechwood Street'\'', '\''Helena, MT 59601'\'', '\''866 Brook Court'\'', '\''Harrison Township, MI 48045'\'', '\''40 Bayport Street'\'', '\''Virginia Beach, VA 23451'\'', '\''20 Hanover St.'\''",
	"mode": "combined"
}'

# Windows cmd ( https://stackoverflow.com/questions/11834238/curl-post-command-line-on-windows-restful-service ):
curl -X POST -d "{""stringsToMatch"":""'31 Beechwood Street','Helena, MT 59601','866 Brook Court','Harrison Township, MI 48045','40 Bayport Street','Virginia Beach, VA 23451','20 Hanover St.'"",
""stringsToMatchIn"":""'31 Beechwood Street','Helena, MT 59601','866 Brook Court','Harrison Township, MI 48045','40 Bayport Street','Virginia Beach, VA 23451','20 Hanover St.'"",
""mode"":""combined""}" http://localhost:8080/api/v1/requests/


# GET REQUEST
# *Nix terminal:
curl --location --request GET 'http://localhost:8080/api/v1/requests/4027cf00-6ff3-4239-9ec5-2d820b4f93cd/' \
--data-raw ''

# Windows cmd:
curl -X GET http://localhost:8080/api/v1/requests/4027cf00-6ff3-4239-9ec5-2d820b4f93cd/
