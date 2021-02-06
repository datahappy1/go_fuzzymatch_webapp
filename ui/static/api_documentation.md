**Creating a new request**
----
  You create a new request only to initiate the fuzzy matching process. This new request has TTL set to 1 hour. After that, the request is removed from the application request database.

* **URL**
  * when running this app @ localhost: 
    * http://localhost:8080/api/v1/requests/
  * when accessing the deployed app API:
    * http://fuzzster.herokuapp.com:8080/api/v1/requests/

* **Method:**
  
  `POST`
  
*  **URL Params**

   None

* **Data Params**

  * stringsToMatch `string` ( a comma separated list of strings for matching, use single quotes to quote strings that contain a comma )
  * stringsToMatchIn `string` ( a comma separated list of strings to match in, use single quotes to quote strings that contain a comma )
  * mode `string` ( one of : `simple` | `deepDive` | `combined` )

* **Success Response:**
  
  On success, the endpoint returns status code 200 and the RequestId.

  * **Code:** 200 <br />
    **Content:** `{ "RequestId" : "aa3c1207-0ee4-4a50-867d-df8da8b8cf1a" }`
 
* **Error Response:**

  <_Most endpoints will have many ways they can fail. From unauthorized access, to wrongful parameters etc. All of those should be liste d here. It might seem repetitive, but it helps prevent assumptions from being made where they should be._>

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "Log in" }`

  OR

  * **Code:** 422 UNPROCESSABLE ENTRY <br />
    **Content:** `{ error : "Email Invalid" }`

* **Sample Call:**

	Windows cmd:
	`curl -X POST -d "{""stringsToMatch"":""'apple, gmbh','corp'"",""stringsToMatchIn"":""aplle"",""mode"":""combined""}" http://localhost:8080/api/v1/requests/`

	*Nix terminal:
	`curl --location --request POST 'http://localhost:8080/api/v1/requests/' \
	--header 'Content-Type: text/plain' \
	--data-raw '{
		"stringsToMatch": "'\''231 Beechwood Street'\'', '\''Helena, MT 59601'\'', '\''866 Brook Court'\'', '\''Harrison Township, MI 48045'\'', '\''40 Bayport Street'\'', '\''Virginia Beach, VA 23451'\'', '\''20 Hanover St.",
			"stringsToMatchIn": "'\''231 Beechwood Street'\'', '\''Helena, MT 59601'\'', '\''866 Brook Court'\'', '\''Harrison Township, MI 48045'\'', '\''40 Bayport Street'\'', '\''Virginia Beach, VA 23451'\'', '\''20 Hanover St.'\''",
			"mode": "combined"
	}'`

* **Notes:**

  <_This is where all uncertainties, commentary, discussion etc. can go. I recommend timestamping and identifying oneself when leaving comments here._> 

**Getting results request**
----
  The fuzzy matching process is lazy evaluated using a following GET request. 

* **URL**

  * when running this app @ localhost: 
    * http://localhost:8080/api/v1/requests/{requestID}/
  * when accessing the deployed app API:
    * http://fuzzster.herokuapp.com:8080/api/v1/requests/{requestID}/


* **Method:**
  
  `GET`
  
*  **URL Params**

   None

* **Data Params**

   None

* **Success Response:**
  
  <_What should the status code be on success and is there any returned data? This is useful when people need to to know what their callbacks should expect!_>

  * **Code:** 200 <br />
    **Content:** `{ id : 12 }`
 
* **Error Response:**

  <_Most endpoints will have many ways they can fail. From unauthorized access, to wrongful parameters etc. All of those should be liste d here. It might seem repetitive, but it helps prevent assumptions from being made where they should be._>

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{ error : "Log in" }`

  OR

  * **Code:** 422 UNPROCESSABLE ENTRY <br />
    **Content:** `{ error : "Email Invalid" }`

* **Sample Call:**

  <_Just a sample call to your endpoint in a runnable format ($.ajax call or a curl request) - this makes life easier and more predictable._> 

* **Notes:**

  <_This is where all uncertainties, commentary, discussion etc. can go. I recommend timestamping and identifying oneself when leaving comments here._> 