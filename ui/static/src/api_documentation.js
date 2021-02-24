import {ApiDocumentationMarkdownFileLocation} from './config.js';

export function convertMarkdownToHtml(inputText) {
    let converter = new showdown.Converter();
    return converter.makeHtml(inputText)
}

export function updateApiDocumentationDiv(content) {
    let apiDocumentationDivElement = document.getElementById("apiDocumentationDiv");

    apiDocumentationDivElement.innerHTML = content;
}

export async function fetch_api_documentation_markdown() {
    const fetchResult = await fetch(ApiDocumentationMarkdownFileLocation,);

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