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

