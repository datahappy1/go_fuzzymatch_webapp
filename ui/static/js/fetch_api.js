const BaseApiRequestsUrl = 'http://localhost:8080/api/v1/requests/';
const ApiDocumentationMarkdownFileLocation = 'http://localhost:8080/api_documentation.md';

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

function DOMUpdateOnLoadDocumentationError(message) {
    updateLoadDocumentationErrorAlert(message);
    toggleLoadDocumentationErrorAlert("show");
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

    const fetchResult = await fetch(BaseApiRequestsUrl, otherParam);

    if (fetchResult.ok) {
        return await fetchResult.json();
    }

    throw {
        type: 'Error',
        message: fetchResult.statusText,
        data: fetchResult.url,
        code: fetchResult.status,
    };
}

async function _fetch_get_lazy_response_results(requestId) {
    const otherParam = {
        headers: {
            "content-type": "application/json; charset=UTF-8"
        },
        method: "GET"
    };

    const fetchResult = await fetch(BaseApiRequestsUrl + requestId + '/', otherParam);

    if (fetchResult.ok) {
        return await fetchResult.json();
    }

    throw {
        type: 'Error',
        message: fetchResult.statusText,
        data: fetchResult.url,
        code: fetchResult.status,
    };
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

async function _fetch_api_documentation_markdown() {
    const fetchResult = await fetch(ApiDocumentationMarkdownFileLocation, );

    if (fetchResult.ok) {
        return await fetchResult.text();
    }

    throw {
        type: 'Error',
        message: fetchResult.statusText,
        data: fetchResult.url,
        code: fetchResult.status,
    };
}
