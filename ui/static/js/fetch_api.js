const BaseUrl = 'http://localhost:8080/api/v1/requestsd/';

function processBackendServiceError(message) {
    updateBackendServiceErrorAlert(message);
    toggleBackendServiceErrorAlert("show");
}

function processBackendServicePass() {
    toggleBackendServiceErrorAlert("hide");
}

function processBackendServiceFetchingDataStart() {
    toggleSubmitButtonWhileLoadingResults("hide");
    clearResultsTable();
    showResultsTable();
}

function processBackendServiceFetchingDataEnd() {
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
    //https://stackoverflow.com/questions/38235715/fetch-reject-promise-and-catch-the-error-if-status-is-not-ok

    console.log(fetchResult);

    if (fetchResult.ok) {
        return await fetchResult.json();
    }

    const responseError = {
        type: 'Error',
        message: fetchResult.url,
        data: fetchResult.statusText,
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
    const getRequestResult = await fetchResult.json();

    if (fetchResult.ok) {
        return getRequestResult;
    }

    const responseError = {
        type: 'Error',
        message: getRequestResult.message || 'Something went wrong',
        data: getRequestResult.data || '',
        code: getRequestResult.code || '',
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
