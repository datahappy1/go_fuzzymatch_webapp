function _alertSwitcher(action, alertDivName) {
    let alertDivComponent = document.getElementById(alertDivName);

    if (action === "show") {
        alertDivComponent.style.display = "block";
    } else if (action === "hide") {
        alertDivComponent.style.display = "none";
    }
}

function _updateLoadDocumentationErrorAlert(errorMessage) {
    let loadDocumentationErrorAlertComponent = document.getElementById("loadDocumentationErrorAlert");

    loadDocumentationErrorAlertComponent.innerHTML = `Load Documentation error: ${errorMessage}`;
}

function _updateBackendServiceErrorAlert(errorMessage) {
    let backendServiceErrorDivComponent = document.getElementById("backendServiceErrorAlert");

    backendServiceErrorDivComponent.innerHTML = `Backend service error: ${errorMessage}`;
}

function _updateMissingMandatoryComponentsAlert(components) {
    let mandatoryFieldsDivComponent = document.getElementById("mandatoryFieldsNotFilledAlert");

    mandatoryFieldsDivComponent.innerHTML = `Mandatory fields not filled: ${components}`;
}

function _updateCopyToClipboardAlert(components) {
    let copyToClipboardDivComponent = document.getElementById("copyToClipboardErrorAlert");

    copyToClipboardDivComponent.innerHTML = `Copy to clipboard error: ${components}`;
}

export function updateOnBackendServiceError(message) {
    _updateBackendServiceErrorAlert(message);
    _alertSwitcher("show", "backendServiceErrorAlert")
}

export function updateOnLoadDocumentationError(message) {
    _updateLoadDocumentationErrorAlert(message);
    _alertSwitcher("show", "loadDocumentationErrorAlert")
}

export function updateOnInputError(message) {
    _updateMissingMandatoryComponentsAlert(message);
    _alertSwitcher("show", "mandatoryFieldsNotFilledAlert")
}

export function updateOnCopyToClipboardError(message) {
    _updateCopyToClipboardAlert(message);
    _alertSwitcher("show", "copyToClipboardErrorAlert")
}

export function hidePreviousMatchErrorsAlerts() {
    _alertSwitcher("hide", "backendServiceErrorAlert");
    _alertSwitcher("hide", "mandatoryFieldsNotFilledAlert");
}

export function hidePreviousLoadDocumentationErrorsAlerts() {
    _alertSwitcher("hide", "loadDocumentationErrorAlert");
}

export function hidePreviousMatchResultsErrorsAlerts() {
    _alertSwitcher("hide", "copyToClipboardErrorAlert");
}