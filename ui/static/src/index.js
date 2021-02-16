import {hidePreviousErrors} from './errors.js';
import {getInputValidationErrors, DOMUpdateOnInputError} from './validation.js';
import {convertMarkdownToHtml, updateApiDocumentationDiv} from './api_documentation.js';


import {fetch_post_new_request, DOMUpdateOnBackendServiceError, 
    update_results_table_with_fetched_data, DOMUpdateOnLoadDocumentationError, 
    DOMUpdateOnBackendServiceFetchingDataStart, DOMUpdateOnBackendServiceFetchingDataEnd} from './fetch_api.js';


import {toggleLoadDocumentationErrorAlert, } from './dom_manipulation.js';




function highlightElement(elementName) {
    let a = document.getElementsByTagName('a');
    for (let i = 0; i < a.length; i++) {
        a[i].classList.remove('active')
    }
    elementName.classList.add('active');
}

function startSearchButtonHandler() {

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
        DOMUpdateOnBackendServiceFetchingDataEnd();

    }

    createRequestStartFetchingChain().catch();
}

function loadStaticPagesHandler() {
    async function prepareApiDocumentationContent() {
        let ApiDocumentationMarkdownContent = null;
        try {
            ApiDocumentationMarkdownContent = await _fetch_api_documentation_markdown().then();
        } catch (e) {
            DOMUpdateOnLoadDocumentationError(JSON.stringify(e));
            return;
        }

        let ApiDocumentationHtmlContent = convertMarkdownToHtml(ApiDocumentationMarkdownContent);

        updateApiDocumentationDiv(ApiDocumentationHtmlContent);

    }

    prepareApiDocumentationContent().catch();

}


// window.onload = function () {
//     loadStaticPagesHandler();
// };

window.addEventListener('load', (event) => {
    console.log('The page has fully loaded');
    loadStaticPagesHandler();
});

const button = document.getElementById( 'submitButton' );
button.addEventListener( 'click', startSearchButtonHandler );
