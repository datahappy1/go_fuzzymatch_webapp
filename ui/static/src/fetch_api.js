import {updateLoadDocumentationErrorAlert, toggleLoadDocumentationErrorAlert,
    updateBackendServiceErrorAlert, toggleBackendServiceErrorAlert,
    toggleSubmitButtonWhileLoadingResults, clearResultsTable, showResultsTable} from './dom_manipulation.js';

const BaseApiRequestsUrl = 'http://localhost:8080/api/v1/requests/';
const ApiDocumentationMarkdownFileLocation = 'http://localhost:8080/api_documentation.md';


function updateLoadDocumentationErrorAlert(errorMessage) {
    let loadDocumentationErrorAlertComponent = document.getElementById("loadDocumentationErrorAlert");

    loadDocumentationErrorAlertComponent.innerHTML = `Load Documentation error: ${errorMessage}`;
    loadDocumentationErrorAlertComponent.style.display = "block";
}

function updateBackendServiceErrorAlert(errorMessage) {
    let backendServiceErrorDivComponent = document.getElementById("backendServiceErrorAlert");

    backendServiceErrorDivComponent.innerHTML = `Backend service error: ${errorMessage}`;
    backendServiceErrorDivComponent.style.display = "block";
}

function clearResultsTable() {
    let container = document.getElementById('resultsTableBody');

    container.innerHTML = '';
}

function showResultsTable() {
    let resultsDivElement = document.getElementById("resultsDiv");

    resultsDivElement.style.display = "block";
}

function jumpToAnchor(anchor) {
    window.location.href = `#${anchor}`;
}


export function DOMUpdateOnBackendServiceError(message) {
    updateBackendServiceErrorAlert(message);
    toggleBackendServiceErrorAlert("show");
    toggleSubmitButtonWhileLoadingResults("show");
}

export function DOMUpdateOnBackendServiceFetchingDataStart() {
    toggleSubmitButtonWhileLoadingResults("hide");
    clearResultsTable();
    showResultsTable();
}

export function DOMUpdateOnBackendServiceFetchingDataEnd() {
    toggleSubmitButtonWhileLoadingResults("show");
    jumpToAnchor("results");
}

export function DOMUpdateOnLoadDocumentationError(message) {
    updateLoadDocumentationErrorAlert(message);
    toggleLoadDocumentationErrorAlert("show");
}

export async function fetch_post_new_request() {
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

export async function fetch_get_lazy_response_results(requestId) {
    const otherParam = {
        headers: {
            "content-type": "application/json; charset=UTF-8"
        },
        method: "GET"
    };

    const fetchResult = await fetch(`${BaseApiRequestsUrl}${requestId}/`, otherParam);

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

export async function update_results_table_with_fetched_data(requestId) {
    let results = await fetch_get_lazy_response_results(requestId);

    if (results["ReturnedAllRows"] === true) {
        updateResultsTable(results["Results"]);

    } else {
        updateResultsTable(results["Results"]);
        await update_results_table_with_fetched_data(requestId);
    }
}

export async function fetch_api_documentation_markdown() {
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

