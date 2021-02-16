export function copyResultsTableToClipboard() {
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

export function downloadResultsTableAsCsv(filename) {
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

export function clearTextarea(textareaName) {
    let textareaElement = document.getElementById(textareaName)

    textareaElement.value = "";
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
