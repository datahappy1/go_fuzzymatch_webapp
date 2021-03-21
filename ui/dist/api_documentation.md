***Creating a new request***
----

  You create a new request only to initiate the fuzzy matching process. This new request has TTL set to 10 minutes. After that, the request is removed from the application request in-memory database.

* **URL** 
    {root_api_url}

* **Method:**
  
    `POST`
  
*  **URL Params**

     None

* **Data Params**

    * stringsToMatch `string` ( a comma separated list of strings for matching, use double quotes to quote strings that contain a comma )
    * stringsToMatchIn `string` ( a comma separated list of strings to match in, use double quotes to quote strings that contain a comma )
    * mode `string` ( one of : `simple` | `deepDive` | `combined` )

* **Success Response:**
  
    On success, the endpoint returns status code 200 and the RequestId.
    Response structure:

    ```json
        {
          "RequestID": "string"
        }
    ```

    * **Code:** 200
    * **Content:** 
      ```json
      {"RequestId" : "0f17955c-1fdd-4bfe-8c66-df8a432f1810"}
      ```

* **Error Response:**

    * **Code:** 429 StatusTooManyRequests
    * **Content:** 
      ```json
      {"error" : "too many overall requests in flight, try later"}
      ```

    OR

    * **Code:** 406 StatusNotAcceptable
    * **Content:**  
      ```json
      {"error" : "cannot read request body"}
      ```

    OR

    * **Code:** 422 StatusUnprocessableEntity
    * **Content:** 
      ```json
      {"error" : "error decoding request data"}
      ```
  
    OR

    * **Code:** 422 StatusUnprocessableEntity
    * **Content:** 
      ```json
      {"error" : "error invalid request"}
      ```

    OR

    * **Code:** 500 StatusInternalServerError
    * **Content:** 
      ```json
      {"error" : "error cannot persist request {request ID}"}
      ```

* **Sample Call:**

  	Windows cmd:

    `
    curl -g -H "Content-type: application/json ; charset=UTF-8" -X POST -d "{\"stringsToMatch\":\"Ellerker,Conry,\\\"Konzelmann, O'Ryan\\\",Dibdin,Audibert,Merrydew\",\"stringsToMatchIn\":\"Mingotti,Tyzack,Maylin,Guiton,Selley,Ferrelli,Rutley,Owthwaite,Liggett\",\"mode\":\"combined\"}" http://localhost:8080/api/v1/requests/
    `

    Linux terminal:

    `
    curl --location --request POST '{root_api_url}' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "stringsToMatch": "Ellerker,Conry,\"Konzelmann, O'\''Ryan\",Dibdin,Audibert,Merrydew",
    "stringsToMatchIn": "Mingotti,Tyzack,Maylin,Guiton,Selley,Ferrelli,Rutley,Owthwaite,Liggett",
    "mode":"combined"
    }'
    `

* **Notes:**


***Getting results request***
----
  The fuzzy matching process is lazy evaluated using a following GET request. Keep polling with this GET request until the flag `ReturnedAllRows` evaluates to true. At that point, all results were returned.

* **URL**

    {root_api_url}/{requestID}/

* **Method:**
  
    `GET`
  
*  **URL Params**

    None

* **Data Params**

    None

* **Success Response:**

    On success, the endpoint returns status code 200 and the response with the fuzzy matching results.
    Response structure:

    ```json
        {
          "RequestID": "string",
          "Mode": "string",
          "RequestedOn": "string",
          "ReturnedAllRows": "bool",
          "Results": [
            {
              "StringToMatch": "string",
              "StringMatched": "string",
              "Result": "int"
            }
          ],
        }
    ```

    * **Code:** 200
    * **Content:** 
      ```json
      {
        "RequestID":"0f17955c-1fdd-4bfe-8c66-df8a432f1810",
        "Mode":"combined",
        "RequestedOn":"2021-03-18T22:39:02",
        "ReturnedAllRows":true,
        "Results":[
            {
              "StringToMatch":"Ellerker",
              "StringMatched":"Selley",
              "Result":57
            },
            {
              "StringToMatch":"Conry",
              "StringMatched":"Guiton",
              "Result":36
            },
            {
              "StringToMatch":"\\Konzelmann, O'Ryan\\",
              "StringMatched":"Tyzack",
              "Result":40
            },
            {
              "StringToMatch":"Dibdin",
              "StringMatched":"Maylin",
              "Result":33
            },
            {
              "StringToMatch":"Audibert",
              "StringMatched":"Guiton",
              "Result":42
            },
            {
              "StringToMatch":"Merrydew",
              "StringMatched":"Ferrelli",
              "Result":50
            }
        ]
      }
      ```

* **Error Response:**

  * **Code:** 406 StatusNotAcceptable
  * **Content:** 
    ```json
    {"error":"need a valid UUID for request ID"}
    ```

  OR

  * **Code:** 404 StatusNotFound
  * **Content:** 
    ```json
    {"error":"request not found"}
    ```

  OR

  * **Code:** 500 StatusInternalServerError
  * **Content:** 
    ```json
    {"error":"error cannot process request {request ID}"}
    ```

* **Sample Call:**

    Linux terminal:

    `
    curl --location --request GET '{root_api_url}/0f17955c-1fdd-4bfe-8c66-df8a432f1810/'
    `

    Windows cmd:

    `
    curl -X GET {root_api_url}/0f17955c-1fdd-4bfe-8c66-df8a432f1810/
    `

* **Notes:**
