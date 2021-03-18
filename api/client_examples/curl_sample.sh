#!/bin/bash

# POST REQUEST
# *Nix terminal:
curl --location --request POST 'http://localhost:8080/api/v1/requests/' \
--header 'Content-Type: application/json' \
--data-raw '{
"stringsToMatch": "Ellerker,Conry,\"Konzelmann O'\''Ryan\",Dibdin,Audibert,Merrydew",
"stringsToMatchIn": "Mingotti,Tyzack,Maylin,Guiton,Selley,Ferrelli,Rutley,Owthwaite,Liggett",
"mode":"combined"
}'

# Windows cmd ( https://stackoverflow.com/questions/11834238/curl-post-command-line-on-windows-restful-service ):
curl -X POST -d "{\"stringsToMatch\":\"Ellerker,Conry,\\Konzelmann O'Ryan\\,Dibdin,Audibert,Merrydew\",\"stringsToMatchIn\":\"Mingotti,Tyzack,Maylin,Guiton,Selley,Ferrelli,Rutley,Owthwaite,Liggett\",\"mode\":\"combined\"}" http://localhost:8080/api/v1/requests/
# or
curl -i -X POST -H "Content-Type: application/json" -d "{""stringsToMatch"":""Ellerker,Conry,Konzelmann O'Ryan,Dibdin,Audibert,Merrydew"",""stringsToMatchIn"":""Mingotti,Tyzack,Maylin,Guiton,Selley,Ferrelli,Rutley,Owthwaite,Liggett"",""mode"":""combined""}" http://localhost:8080/api/v1/requests/

# GET REQUEST
# *Nix terminal:
curl --location --request GET 'http://localhost:8080/api/v1/requests/4027cf00-6ff3-4239-9ec5-2d820b4f93cd/'

# Windows cmd:
curl -X GET http://localhost:8080/api/v1/requests/4027cf00-6ff3-4239-9ec5-2d820b4f93cd/
