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

function _getRangeInputSliderValue() {
    let sliderElement = document.getElementById("rangeInput");

    return sliderElement.value
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

export function filterResultsTable() {
    let inputValue, table, tr, td, i, cellValue;

    inputValue = _getRangeInputSliderValue();

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

export function clearResultsTable() {
    let container = document.getElementById('resultsTableBody');

    container.innerHTML = '';
}

export function showResultsDiv() {
    let resultsDivElement = document.getElementById("results");

    resultsDivElement.style.display = "block";
}

export function toggleSubmitButtonWhileLoadingResults(action) {
    let submitButtonElement = document.getElementById("submitButton");
    let submitButtonSpinnerElement = document.getElementById("submitButtonSpinner");

    if (action === "show") {
        submitButtonElement.style.display = "block";
        submitButtonSpinnerElement.style.display = "none";
    } else if (action === "hide") {
        submitButtonElement.style.display = "none";
        submitButtonSpinnerElement.style.display = "block";
    }
}

export function jumpToAnchor(anchor) {
    window.location.href = `#${anchor}`;
}
