const BaseUrl = 'http://localhost:8080/api/v1/requests/';

function createRequestStartFetchingChain() {
    let inputStringsToMatch = document.getElementById("stringsToMatch").value;
    let inputStringsToMatchIn = document.getElementById("stringsToMatchIn").value;
    let inputMode = document.getElementById("mode").value;

    console.log(inputStringsToMatch, inputStringsToMatchIn, inputMode)

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
    const otherParam2 = {
        headers: {
            "content-type": "application/json; charset=UTF-8"
        },
        method: "GET"
    };

    fetch(BaseUrl + requestId + '/', otherParam2)
        .then(data => {
            return data.json()
        })
        .then(res => {
            console.log(res)
            if (res["ReturnedAllRows"] === true) {
                updateResultsTable(res["Results"])
            } else {
                updateResultsTable(res["Results"]);
                fetchResults(requestId);
            }
        })
        .catch(error => console.log(error))

}
