function copyResultsTableToClipboard() {
    let clipboard = new ClipboardJS('.btn');

    clipboard.on('success', function (e) {
        console.info('Action:', e.action);
        console.info('Text:', e.text);
        console.info('Trigger:', e.trigger);

        e.clearSelection();
    });

    clipboard.on('error', function (e) {
        console.error('Action:', e.action);
        console.error('Trigger:', e.trigger);
    });
}

function downloadCSV(csv, filename) {
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

function downloadResultsTableAsCsv(filename) {
    let csv = [];
    let rows = document.querySelectorAll("table tr");

    for (let i = 0; i < rows.length; i++) {
        let row = [], cols = rows[i].querySelectorAll("td, th");

        for (let j = 0; j < cols.length; j++)
            row.push(cols[j].innerText);

        csv.push(row.join(","));
    }

    downloadCSV(csv.join("\n"), filename);
}

function clearTextarea(textareaname) {
    let textareaElement = document.getElementById(textareaname)

    textareaElement.value = "";
}

function clearResultsTable() {
    let container = document.getElementById('resultsTableBody');

    container.innerHTML = '';
}

function updateResultsTable(results) {
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

function getRangeInputSliderValue() {
    let sliderElement = document.getElementById("rangeInput");

    return sliderElement.value
}

function filterResultsTable() {
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

function updateMissingMandatoryComponentsAlert(components) {
    let mandatoryFieldsDivComponent = document.getElementById("mandatoryFieldsNotFilledAlert");

    mandatoryFieldsDivComponent.innerHTML = "Mandatory fields not filled: " + components;

}

function toggleMissingMandatoryComponentsAlert(action) {
    let mandatoryFieldsDivComponent = document.getElementById("mandatoryFieldsNotFilledAlert");

    if (action === "show") {
        mandatoryFieldsDivComponent.style.display = "block";
    } else if (action === "hide") {
        mandatoryFieldsDivComponent.style.display = "none";
    }
}

function toggleSubmitButtonWhileLoadingResults(action) {
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

function updateBackendServiceErrorAlert(errorMessage) {
    let backendServiceErrorDivComponent = document.getElementById("backendServiceErrorAlert");

    backendServiceErrorDivComponent.innerHTML = "Backend service error: " + errorMessage;
    backendServiceErrorDivComponent.style.display = "block";
}

function toggleBackendServiceErrorAlert(action) {
    let backendServiceErrorDivComponent = document.getElementById("backendServiceErrorAlert");

    if (action === "show") {
        backendServiceErrorDivComponent.style.display = "block";
    } else if (action === "hide") {
        backendServiceErrorDivComponent.style.display = "none";
    }
}

function updateLoadDocumentationErrorAlert(errorMessage) {
    let loadDocumentationErrorAlertComponent = document.getElementById("loadDocumentationErrorAlert");

    loadDocumentationErrorAlertComponent.innerHTML = "Load Documentation error: " + errorMessage;
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

function jumpToAnchor(anchor) {
    window.location.href = "#" + anchor;
}

function showResultsTable() {
    let resultsDivElement = document.getElementById("resultsDiv");

    resultsDivElement.style.display = "block";
}

function hidePreviousErrors() {
    toggleMissingMandatoryComponentsAlert("hide");
    toggleBackendServiceErrorAlert("hide");
    toggleSubmitButtonWhileLoadingResults("show");
}

function convertMarkdownToHtml(inputText) {
    let converter = new showdown.Converter();
    let html = converter.makeHtml(inputText);

    return html
}

function updateApiDocumentationDiv(content) {
    let apiDocumentationDivElement = document.getElementById("apiDocumentationDiv");

    apiDocumentationDivElement.innerHTML = content;
}


function highlightElement(elementName) {
    var a = document.getElementsByTagName('a');
    for (i = 0; i < a.length; i++) {
        a[i].classList.remove('active')
    }
    elementName.classList.add('active');
}