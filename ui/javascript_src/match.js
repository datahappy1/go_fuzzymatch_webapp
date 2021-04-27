import {BaseApiRequestsUrl} from './config.js';

function _updateResultsTable(results) {
    let htmlResultsTable = '';
    let container = document.getElementById('resultsTableBody');

    for (let i = 0; i < results.length; i++) {
        let htmlTableRow = `<tr>
                        <td>${results[i]["StringToMatch"]}</td>
                        <td>${results[i]["StringMatched"]}</td>
                        <td>${results[i]["Result"]}</td>
                    </tr>`;

        htmlResultsTable += htmlTableRow;
    }

    container.innerHTML += htmlResultsTable;
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
    const url = BaseApiRequestsUrl;
    let fetchResult = null;

    try {
        fetchResult = await fetch(url, otherParam);
    } catch (e) {
        throw {
            type: 'Error',
            message: e.message,
            data: url,
            code: 500,
        };
    }

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
    const url = `${BaseApiRequestsUrl}${requestId}/`;
    let fetchResult = null;

    try {
        fetchResult = await fetch(url, otherParam);
    } catch (e) {
        throw {
            type: 'Error',
            message: e.message,
            data: url,
            code: 500,
        };
    }

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

    _updateResultsTable(results["Results"]);

    if (results["ReturnedAllRows"] === false) {
        await update_results_table_with_fetched_data(requestId);
    }
}

export function clearStringsTextarea(elementName) {
    let TextAreaElement = document.getElementById(elementName);
    TextAreaElement.value = "";
}
