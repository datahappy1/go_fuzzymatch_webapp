const BaseUrl = 'http://localhost:8080/api/v1/requests/';

function DOMUpdateOnBackendServiceError(message) {
    updateBackendServiceErrorAlert(message);
    toggleBackendServiceErrorAlert("show");
    toggleSubmitButtonWhileLoadingResults("show");
}

function DOMUpdateOnBackendServiceFetchingDataStart() {
    toggleSubmitButtonWhileLoadingResults("hide");
    clearResultsTable();
    showResultsTable();
}

function DOMUpdateOnBackendServiceFetchingDataEnd() {
    toggleSubmitButtonWhileLoadingResults("show");
    jumpToAnchor("results");
}

async function _fetch_post_new_request() {
    const inputStringsToMatch = document.getElementById("stringsToMatch").value;
    const inputStringsToMatchIn = document.getElementById("stringsToMatchIn").value;
    const inputMode = document.getElementById("mode").value;

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

    const fetchResult = await fetch(BaseUrl, otherParam);

    if (fetchResult.ok) {
        return await fetchResult.json();
    }

    const responseError = {
        type: 'Error',
        message: fetchResult.statusText,
        data: fetchResult.url,
        code: fetchResult.status,
    };

    let error = new Error();
    error = {...error, ...responseError};
    throw (error);
}

async function _fetch_get_lazy_response_results(requestId) {
    const otherParam = {
        headers: {
            "content-type": "application/json; charset=UTF-8"
        },
        method: "GET"
    };

    const fetchResult = await fetch(BaseUrl + requestId + '/', otherParam);

    if (fetchResult.ok) {
        return await fetchResult.json();
    }

    const responseError = {
        type: 'Error',
        message: fetchResult.statusText,
        data: fetchResult.url,
        code: fetchResult.status,
    };

    let error = new Error();
    error = {...error, ...responseError};
    throw (error);
}

async function _update_results_table_with_fetched_data(requestId) {
    let results = await _fetch_get_lazy_response_results(requestId);

    if (results["ReturnedAllRows"] === true) {
        updateResultsTable(results["Results"]);

    } else {
        updateResultsTable(results["Results"]);
        await _update_results_table_with_fetched_data(requestId);
    }
}
