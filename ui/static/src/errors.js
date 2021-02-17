function _toggleMissingMandatoryComponentsAlert(action) {
    let mandatoryFieldsDivComponent = document.getElementById("mandatoryFieldsNotFilledAlert");

    if (action === "show") {
        mandatoryFieldsDivComponent.style.display = "block";
    } else if (action === "hide") {
        mandatoryFieldsDivComponent.style.display = "none";
    }
}

function _toggleBackendServiceErrorAlert(action) {
    let backendServiceErrorDivComponent = document.getElementById("backendServiceErrorAlert");

    if (action === "show") {
        backendServiceErrorDivComponent.style.display = "block";
    } else if (action === "hide") {
        backendServiceErrorDivComponent.style.display = "none";
    }
}

function _updateLoadDocumentationErrorAlert(errorMessage) {
    let loadDocumentationErrorAlertComponent = document.getElementById("loadDocumentationErrorAlert");

    loadDocumentationErrorAlertComponent.innerHTML = `Load Documentation error: ${errorMessage}`;
    loadDocumentationErrorAlertComponent.style.display = "block";
}

function _updateBackendServiceErrorAlert(errorMessage) {
    let backendServiceErrorDivComponent = document.getElementById("backendServiceErrorAlert");

    backendServiceErrorDivComponent.innerHTML = `Backend service error: ${errorMessage}`;
    backendServiceErrorDivComponent.style.display = "block";
}

function _updateMissingMandatoryComponentsAlert(components) {
    let mandatoryFieldsDivComponent = document.getElementById("mandatoryFieldsNotFilledAlert");

    mandatoryFieldsDivComponent.innerHTML = `Mandatory fields not filled: ${components}`;
}

export function DOMUpdateOnBackendServiceError(message) {
    _updateBackendServiceErrorAlert(message);
    toggleBackendServiceErrorAlert("show");
    toggleSubmitButtonWhileLoadingResults("show");
}

export function DOMUpdateOnLoadDocumentationError(message) {
    _updateLoadDocumentationErrorAlert(message);
    _toggleLoadDocumentationErrorAlert("show");
}

export function DOMUpdateOnInputError(message) {
    _updateMissingMandatoryComponentsAlert(message);
    _toggleMissingMandatoryComponentsAlert("show");
}

export function hidePreviousErrors() {
    _toggleMissingMandatoryComponentsAlert("hide");
    _toggleBackendServiceErrorAlert("hide");
}
