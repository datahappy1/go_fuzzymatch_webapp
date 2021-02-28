import * as ShowdownJS from "./external/showdown.js/1.9.1/showdown.js";

import { ApiDocumentationMarkdownFileLocation } from './config.js';

export function convertMarkdownToHtml(inputText) {
    let converter = new ShowdownJS.Converter();
    return converter.makeHtml(inputText)
}

export function updateApiDocumentationDiv(content) {
    let apiDocumentationDivElement = document.getElementById("apiDocumentationDiv");

    apiDocumentationDivElement.innerHTML = content;
}

export async function fetch_api_documentation_markdown() {
    let fetchResult = null;
    try {
        fetchResult = await fetch(ApiDocumentationMarkdownFileLocation);
    }
    catch (e) {
        throw {
            type: 'Error',
            message: e.message,
            data: ApiDocumentationMarkdownFileLocation,
            code: 500,
        };
    }
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
