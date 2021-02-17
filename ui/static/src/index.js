import {
    DOMUpdateOnBackendServiceError,
    DOMUpdateOnBackendServiceFetchingDataEnd,
    DOMUpdateOnBackendServiceFetchingDataStart,
    DOMUpdateOnInputError,
    DOMUpdateOnLoadDocumentationError,
    hidePreviousErrors
} from "./errors.js";
import {getInputValidationErrors} from './validation.js';
import {
    convertMarkdownToHtml,
    fetch_api_documentation_markdown,
    updateApiDocumentationDiv
} from './api_documentation.js';
import {fetch_post_new_request, update_results_table_with_fetched_data} from "./match.js";
import {
    showResultsTable,
    clearResultsTable,
    copyResultsTableToClipboard,
    downloadResultsTableAsCsv,
    getRangeInputSliderValue
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

// function highlightElementHandler(elementName) {
//     let a = document.getElementsByTagName('a');
//     for (let i = 0; i < a.length; i++) {
//         a[i].classList.remove('active')
//     }
//     elementName.classList.add('active');
// }

function startMatchButtonHandler() {

    async function createRequestStartFetchingChain() {
        hidePreviousErrors();

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

        DOMUpdateOnBackendServiceFetchingDataStart();
        try {
            await update_results_table_with_fetched_data(requestId);
        } catch (e) {
            DOMUpdateOnBackendServiceError(JSON.stringify(e));
            return;
        }
        //clearResultsTable();
        showResultsTable();
        DOMUpdateOnBackendServiceFetchingDataEnd();

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

// window.onload = function () {
//     loadStaticPagesHandler();
// };
window.addEventListener('load', (event) => {
    console.log('The page has fully loaded');
    loadStaticPagesHandler();
});

// const highlightElementLink = document.getElementById('submitButton');
// highlightElementLink.addEventListener('click', highlightElementHandler);

const matchButton = document.getElementById('submitButton');
matchButton.addEventListener('click', startMatchButtonHandler);

const filterResultsButton = document.getElementById('rangeInput');
filterResultsButton.addEventListener('onchange', filterResultsTableButtonHandler);

const copyResultsToClipboardButton = document.getElementById('copyResultsTableToClipboardButton');
copyResultsToClipboardButton.addEventListener('onchange', copyResultsToClipboardButtonHandler);

const downloadResultsToCSVButton = document.getElementById('downloadResultsTableAsCsvButton');
downloadResultsToCSVButton.addEventListener('onchange', downloadResultsCSVButtonHandler);

const clearResultsTableButton = document.getElementById('clearResultsTableButton');
clearResultsTableButton.addEventListener('onchange', clearResultsButtonHandler);




