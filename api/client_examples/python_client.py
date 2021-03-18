import json

import requests

requests_url = "http://localhost:8080/api/v1/requests/"

payload = json.dumps({"stringsToMatch": "Ellerker,Conry,\"Konzelmann O'Ryan\",Dibdin,Audibert,Merrydew",
                      "stringsToMatchIn": "Mingotti,Tyzack,Maylin,Guiton,Selley,Ferrelli,Rutley,Owthwaite,Liggett",
                      "mode": "combined"})

headers = {
  'Content-Type': 'application/json'
}

post_request_response = requests.request("POST", requests_url, headers=headers, data=payload)

if post_request_response.status_code == 200:
    post_response_json = json.loads(post_request_response.text)

    request_id = post_response_json["RequestID"]

    request_url = f"http://localhost:8080/api/v1/requests/{request_id}/"

    payload = {}
    headers = {}

    while True:
        get_request_response = requests.request("GET", request_url, headers=headers, data=payload)
        get_response_json = json.loads(get_request_response.text)

        print(get_response_json)

        if get_response_json["ReturnedAllRows"] is True:
            break
