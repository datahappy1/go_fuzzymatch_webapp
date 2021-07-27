import {
    updateOnBackendServiceError,
    updateOnCopyToClipboardError,
    updateOnInputError,
    updateOnLoadDocumentationError,
    hidePreviousMatchErrorsAlerts,
    hidePreviousMatchResultsErrorsAlerts,
    hidePreviousLoadDocumentationErrorsAlerts
} from "./errors.js";
import {highlightActiveMenuItem} from './menu.js';
import {getInputValidationErrors} from './validation.js';
import {
    convertMarkdownToHtml,
    fetch_api_documentation_markdown,
    updateApiDocumentationDiv
} from './api_documentation.js';
import {
    clearStringsTextarea,
    fetchPostNewRequest,
    updateResultsTableWithFetchedData
} from "./match.js";
import {
    clearResultsTable,
    copyResultsTableToClipboard,
    downloadResultsTableAsCsv,
    jumpToAnchor,
    showResultsDiv,
    toggleSubmitButtonWhileLoadingResults
} from "./match_results.js";
import {filterResultsTable} from "./match_results";


function loadStaticPagesHandler() {
    async function prepareApiDocumentationContent() {
        hidePreviousLoadDocumentationErrorsAlerts();

        let ApiDocumentationMarkdownContent = null;
        try {
            ApiDocumentationMarkdownContent = await fetch_api_documentation_markdown().then();
        } catch (e) {
            updateOnLoadDocumentationError(JSON.stringify(e));
            return;
        }

        let ApiDocumentationHtmlContent = convertMarkdownToHtml(ApiDocumentationMarkdownContent);

        updateApiDocumentationDiv(ApiDocumentationHtmlContent);

    }

    prepareApiDocumentationContent().then();

}

function startMatchButtonHandler() {

    async function createRequestStartFetchingChain() {
        hidePreviousMatchErrorsAlerts();
        toggleSubmitButtonWhileLoadingResults("show");

        let inputValidationErrors = getInputValidationErrors();
        if (inputValidationErrors.length > 0) {
            updateOnInputError(inputValidationErrors);
            return;
        }

        let requestId = null;
        try {
            const objectId = await fetchPostNewRequest();
            requestId = objectId["RequestID"];
        } catch (e) {
            updateOnBackendServiceError(JSON.stringify(e));
            toggleSubmitButtonWhileLoadingResults("show");
            return;
        }

        toggleSubmitButtonWhileLoadingResults("hide");
        clearResultsTable();
        showResultsDiv();

        try {
            await updateResultsTableWithFetchedData(requestId);
        } catch (e) {
            updateOnBackendServiceError(JSON.stringify(e));
            toggleSubmitButtonWhileLoadingResults("show");
            return;
        }
        toggleSubmitButtonWhileLoadingResults("show");
        jumpToAnchor("results");
    }

    createRequestStartFetchingChain().then();

}

function filterResultsTableButtonHandler() {
    filterResultsTable();
}

function copyResultsToClipboardButtonHandler() {
    hidePreviousMatchResultsErrorsAlerts();
    try {
        copyResultsTableToClipboard();
    } catch (e) {
        updateOnCopyToClipboardError(JSON.stringify(e));
    }
}

function downloadResultsCSVButtonHandler() {
    downloadResultsTableAsCsv('fuzzy_match_results.csv')
}

function clearResultsButtonHandler() {
    clearResultsTable()
}

function clearStringsToMatchTextareaHandler() {
    clearStringsTextarea("stringsToMatch");
}

function clearStringsToMatchInTextareaHandler() {
    clearStringsTextarea("stringsToMatchIn");
}

window.addEventListener('load', (event) => {
    highlightActiveMenuItem();
    loadStaticPagesHandler();
});

const matchButton = document.getElementById('submitButton');
matchButton.addEventListener('click', startMatchButtonHandler);

const clearStringsToMatchTextarea = document.getElementById('clearStringsToMatchTextarea');
clearStringsToMatchTextarea.addEventListener('click', clearStringsToMatchTextareaHandler);

const clearStringsToMatchInTextarea = document.getElementById('clearStringsToMatchInTextarea');
clearStringsToMatchInTextarea.addEventListener('click', clearStringsToMatchInTextareaHandler);

const filterResultsButton = document.getElementById('rangeInput');
filterResultsButton.addEventListener('change', filterResultsTableButtonHandler);

const copyResultsToClipboardButton = document.getElementById('copyResultsTableToClipboardButton');
copyResultsToClipboardButton.addEventListener('click', copyResultsToClipboardButtonHandler);

const downloadResultsToCSVButton = document.getElementById('downloadResultsTableAsCsvButton');
downloadResultsToCSVButton.addEventListener('click', downloadResultsCSVButtonHandler);

const clearResultsTableButton = document.getElementById('clearResultsTableButton');
clearResultsTableButton.addEventListener('click', clearResultsButtonHandler);
