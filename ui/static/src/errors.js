function _toggleLoadDocumentationErrorAlert(action) {
    let loadDocumentationErrorAlertComponent = document.getElementById("loadDocumentationErrorAlert");

    if (action === "show") {
        loadDocumentationErrorAlertComponent.style.display = "block";
    } else if (action === "hide") {
        loadDocumentationErrorAlertComponent.style.display = "none";
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

// doesnt belong here TODO
function _toggleSubmitButtonWhileLoadingResults(action) {
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

export function hidePreviousErrors() {
    _toggleMissingMandatoryComponentsAlert("hide");
    _toggleBackendServiceErrorAlert("hide");
    _toggleSubmitButtonWhileLoadingResults("show");
}
