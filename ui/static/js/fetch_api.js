const BaseUrl = 'http://localhost:8080/api/v1/requests/';

function createRequestStartFetchingChain() {
    const inputStringsToMatch = document.getElementById("stringsToMatch").value;
    const inputStringsToMatchIn = document.getElementById("stringsToMatchIn").value;
    const inputMode = document.getElementById("mode").value;

    //console.log(inputStringsToMatch, inputStringsToMatchIn, inputMode)

    const Data = JSON.stringify({
        stringsToMatch: inputStringsToMatch,
        stringsToMatchIn: inputStringsToMatchIn,
        mode: inputMode
    });
    const otherParam = {
        headers: {
            "content-type": "application/json; charset=UTF-8"
        },
        body: Data,
        method: "POST"
    };

    fetch(BaseUrl, otherParam)
        .then(data => {
            return data.json()
        })
        .then(res => {
            fetchResults(res["RequestID"])
        })
        .catch(error => console.log(error))
}

function fetchResults(requestId) {
    const otherParam = {
        headers: {
            "content-type": "application/json; charset=UTF-8"
        },
        method: "GET"
    };

    fetch(BaseUrl + requestId + '/', otherParam)
        .then(data => {
            return data.json()
        })
        .then(res => {
            // console.log(res);
            //console.log(res["Results"].length / inputStringsToMatch.length)
            //updateProgressBar(res["Results"].length / inputStringsToMatch.length);

            if (res["ReturnedAllRows"] === true) {
                updateResultsTable(res["Results"]);
            } else {
                updateResultsTable(res["Results"]);
                fetchResults(requestId);
            }
        })
        .catch(error => console.log(error))

}
