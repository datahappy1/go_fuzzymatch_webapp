import {
    DOMUpdateOnBackendServiceError,
    DOMUpdateOnInputError,
    DOMUpdateOnLoadDocumentationError,
    hidePreviousErrors
} from "./errors.js";
import { getInputValidationErrors }
    from './validation.js';
import {
    convertMarkdownToHtml,
    fetch_api_documentation_markdown,
    updateApiDocumentationDiv
} from './api_documentation.js';
import {
    fetch_post_new_request,
    update_results_table_with_fetched_data,
} from "./match.js";
import {
    showResultsTable,
    clearResultsTable,
    copyResultsTableToClipboard,
    downloadResultsTableAsCsv,
    getRangeInputSliderValue,
    toggleSubmitButtonWhileLoadingResults,
    jumpToAnchor
} from "./match_results.js";


function loadStaticPagesHandler() {
    async function prepareApiDocumentationContent() {
        let ApiDocumentationMarkdownContent = null;
        try {
            ApiDocumentationMarkdownContent = await fetch_api_documentation_markdown().then();
        } catch (e) {
            DOMUpdateOnLoadDocumentationError(JSON.stringify(e));
            return;
        }

        let ApiDocumentationHtmlContent = convertMarkdownToHtml(ApiDocumentationMarkdownContent);

        updateApiDocumentationDiv(ApiDocumentationHtmlContent);

    }

    prepareApiDocumentationContent().catch();

}

function startMatchButtonHandler() {

    async function createRequestStartFetchingChain() {
        hidePreviousErrors();
        toggleSubmitButtonWhileLoadingResults("show");

        let inputValidationErrors = getInputValidationErrors();
        if (inputValidationErrors.length > 0) {
            DOMUpdateOnInputError(inputValidationErrors);
            return;
        }

        let requestId = null;
        try {
            const objectId = await fetch_post_new_request();
            requestId = objectId["RequestID"];
        } catch (e) {
            DOMUpdateOnBackendServiceError(JSON.stringify(e));
            return;
        }

        toggleSubmitButtonWhileLoadingResults("hide");
        clearResultsTable();
        try {
            await update_results_table_with_fetched_data(requestId);
        } catch (e) {
            DOMUpdateOnBackendServiceError(JSON.stringify(e));
            return;
        }
        showResultsTable();
        toggleSubmitButtonWhileLoadingResults("show");
        jumpToAnchor("results");

    }

    createRequestStartFetchingChain().catch();
}

function filterResultsTableButtonHandler() {
    let inputValue, table, tr, td, i, cellValue;

    inputValue = getRangeInputSliderValue();

    table = document.getElementById("resultsTable");
    tr = table.getElementsByTagName("tr");

    for (i = 0; i < tr.length; i++) {
        td = tr[i].getElementsByTagName("td")[2];
        if (td) {
            cellValue = td.textContent || td.innerText;
            if (+cellValue >= +inputValue) {
                tr[i].style.display = "";
            } else {
                tr[i].style.display = "none";
            }
        }
    }
}

function copyResultsToClipboardButtonHandler() {
    copyResultsTableToClipboard()
}

function downloadResultsCSVButtonHandler() {
    downloadResultsTableAsCsv('fuzzy_match_results.csv')
}

function clearResultsButtonHandler() {
    clearResultsTable()
}

function clearStringsToMatchTextareaHandler() {
    element = document.getElementById("clearStringsToMatchTextarea")
    element.value = "";
}

function clearStringsToMatchInTextareaHandler() {
    element = document.getElementById("clearStringsToMatchInTextarea")
    element.value = "";
}

window.addEventListener('load', (event) => {
    loadStaticPagesHandler();
});

const matchButton = document.getElementById('submitButton');
matchButton.addEventListener('click', startMatchButtonHandler);

const clearStringsToMatchTextarea = document.getElementById('clearStringsToMatchTextarea');
clearStringsToMatchTextarea.addEventListener('click', clearStringsToMatchTextareaHandler)

const clearStringsToMatchInTextarea = document.getElementById('clearStringsToMatchInTextarea');
clearStringsToMatchInTextarea.addEventListener('click', clearStringsToMatchInTextareaHandler)

const filterResultsButton = document.getElementById('rangeInput');
filterResultsButton.addEventListener('change', filterResultsTableButtonHandler);

const copyResultsToClipboardButton = document.getElementById('copyResultsTableToClipboardButton');
copyResultsToClipboardButton.addEventListener('click', copyResultsToClipboardButtonHandler);

const downloadResultsToCSVButton = document.getElementById('downloadResultsTableAsCsvButton');
downloadResultsToCSVButton.addEventListener('click', downloadResultsCSVButtonHandler);

const clearResultsTableButton = document.getElementById('clearResultsTableButton');
clearResultsTableButton.addEventListener('click', clearResultsButtonHandler);
