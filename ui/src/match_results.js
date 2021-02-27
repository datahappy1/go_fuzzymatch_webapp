import * as ClipboardJS from "./external/clipboard.js/2.0.6/clipboard.js";

function _downloadCSV(csv, filename) {
    let csvFile;
    let downloadLink;

    csvFile = new Blob([csv], {type: "text/csv"});

    downloadLink = document.createElement("a");
    downloadLink.download = filename;
    downloadLink.href = window.URL.createObjectURL(csvFile);
    downloadLink.style.display = "none";

    document.body.appendChild(downloadLink);
    downloadLink.click();
}

export function downloadResultsTableAsCsv(filename) {
    let csv = [];
    let rows = document.querySelectorAll("table tr");

    for (let i = 0; i < rows.length; i++) {
        let row = [], cols = rows[i].querySelectorAll("td, th");

        for (let j = 0; j < cols.length; j++)
            row.push(cols[j].innerText);

        csv.push(row.join(","));
    }

    _downloadCSV(csv.join("\n"), filename);
}

export function copyResultsTableToClipboard() {
    let clipboard = new ClipboardJS('.btn');

    clipboard.on('success', function (e) {
        e.clearSelection();
    });

    clipboard.on('error', function (e) {
        throw {
            type: 'Error',
            message: e.action,
            data: e.trigger,
            code: null,
        };
    });
}

export function getRangeInputSliderValue() {
    let sliderElement = document.getElementById("rangeInput");

    return sliderElement.value
}

export function clearResultsTable() {
    let container = document.getElementById('resultsTableBody');

    container.innerHTML = '';
}

export function showResultsTable() {
    let resultsDivElement = document.getElementById("results");

    resultsDivElement.style.display = "block";
}

export function toggleSubmitButtonWhileLoadingResults(action) {
    let submitButtonElement = document.getElementById("submitButton");
    let submitButtonSpinnerElement = document.getElementById("submitButtonSpinner");

    if (action === "show") {
        submitButtonElement.disabled = false;
        submitButtonSpinnerElement.style.display = "none";
    } else if (action === "hide") {
        submitButtonElement.disabled = true;
        submitButtonSpinnerElement.style.display = "block";
    }
}

export function jumpToAnchor(anchor) {
    window.location.href = `#${anchor}`;
}
