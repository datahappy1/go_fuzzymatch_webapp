const _mandatoryInputComponents = [
    {'componentId': 'stringsToMatch', 'componentNameForAlertMessage': '"strings to match" textarea'},
    {'componentId': 'stringsToMatchIn', 'componentNameForAlertMessage': '"strings to match in" textarea'}
];

function _isValidMandatoryComponent(elementId) {
    let validatedComponent = document.getElementById(elementId);

    return !(validatedComponent.value.length === 0);
}

function _getMissingMandatoryInputs() {
    let missingMandatoryInputComponents = [];

    _mandatoryInputComponents.map(item => {
        if (_isValidMandatoryComponent(item.componentId) === false) {
            missingMandatoryInputComponents.push(item.componentNameForAlertMessage);
        }

    });

    return missingMandatoryInputComponents;
}

function _updateMissingMandatoryComponentsAlert(components) {
    let mandatoryFieldsDivComponent = document.getElementById("mandatoryFieldsNotFilledAlert");

    mandatoryFieldsDivComponent.innerHTML = `Mandatory fields not filled: ${components}`;
}

function _toggleMissingMandatoryComponentsAlert(action) {
    let mandatoryFieldsDivComponent = document.getElementById("mandatoryFieldsNotFilledAlert");

    if (action === "show") {
        mandatoryFieldsDivComponent.style.display = "block";
    } else if (action === "hide") {
        mandatoryFieldsDivComponent.style.display = "none";
    }
}

export function getInputValidationErrors() {
    let missingMandatoryInputs = _getMissingMandatoryInputs();

    if (missingMandatoryInputs.length > 0) {
        return missingMandatoryInputs
    }
    return []
}

export function DOMUpdateOnInputError(message) {
    _updateMissingMandatoryComponentsAlert(message);
    _toggleMissingMandatoryComponentsAlert("show");
}
