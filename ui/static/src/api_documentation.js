export function convertMarkdownToHtml(inputText) {
    let converter = new showdown.Converter();
    let html = converter.makeHtml(inputText);

    return html
}

export function updateApiDocumentationDiv(content) {
    let apiDocumentationDivElement = document.getElementById("apiDocumentationDiv");

    apiDocumentationDivElement.innerHTML = content;
}
