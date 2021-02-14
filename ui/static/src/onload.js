function updateLoadDocumentationErrorAlert(errorMessage) {
    let loadDocumentationErrorAlertComponent = document.getElementById("loadDocumentationErrorAlert");

    loadDocumentationErrorAlertComponent.innerHTML = `Load Documentation error: ${errorMessage}`;
    loadDocumentationErrorAlertComponent.style.display = "block";
}

function toggleLoadDocumentationErrorAlert(action) {
    let loadDocumentationErrorAlertComponent = document.getElementById("loadDocumentationErrorAlert");

    if (action === "show") {
        loadDocumentationErrorAlertComponent.style.display = "block";
    } else if (action === "hide") {
        loadDocumentationErrorAlertComponent.style.display = "none";
    }
}

function DOMUpdateOnLoadDocumentationError(message) {
    updateLoadDocumentationErrorAlert(message);
    toggleLoadDocumentationErrorAlert("show");
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

window.onload = function () {
    loadStaticPagesHandler();
};
