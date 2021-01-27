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

// https://www.w3schools.com/howto/howto_js_filter_table.asp
function filterResultsTable() {
    let input, table, tr, td, i, cellValue;
    input = getRangeInputSliderValue();
    table = document.getElementById("resultsTable");
    tr = table.getElementsByTagName("tr");

    for (i = 0; i < tr.length; i++) {
        td = tr[i].getElementsByTagName("td")[2];
        if (td) {
            cellValue = td.textContent || td.innerText;
            if (cellValue >= input) {
                tr[i].style.display = "";
            } else {
                tr[i].style.display = "none";
            }
        }
    }
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
