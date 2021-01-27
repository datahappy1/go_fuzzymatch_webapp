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

function isValidMandatoryComponents() {
    let stringsToMatchTextarea = document.getElementById("stringsToMatch");
    let stringsToMatchInTextarea = document.getElementById("stringsToMatchIn");

    return !(stringsToMatchTextarea.value.length === 0 || stringsToMatchInTextarea.value.length === 0);
}

function toggleMissingMandatoryComponentsAlert(action) {
    let mandatoryFieldsDivComponent = document.getElementById("mandatoryFieldsNotFilledAlert");
    if (action === "show") {
        mandatoryFieldsDivComponent.style.display = "block";
    } else if (action === "hide") {
        mandatoryFieldsDivComponent.style.display = "none";
    }
}