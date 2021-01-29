const mandatoryInputComponents = [
    {'componentId': 'stringsToMatch', 'componentNameForAlertMessage': '"strings to match" textarea'},
    {'componentId': 'stringsToMatchIn', 'componentNameForAlertMessage': '"strings to match in" textarea'}
];

function isValidMandatoryComponent(elementId) {
    let validatedComponent = document.getElementById(elementId);

    return !(validatedComponent.value.length === 0);
}

function getMissingMandatoryInputs() {
    let missingMandatoryInputComponents = [];

    mandatoryInputComponents.map(item => {
        if (isValidMandatoryComponent(item.componentId) === false) {
            missingMandatoryInputComponents.push(item.componentNameForAlertMessage);
        }

    });

    return missingMandatoryInputComponents;
}

function processInputError(message) {
    updateMissingMandatoryComponentsAlert(message);
    toggleMissingMandatoryComponentsAlert("show");
}

function processInputPass() {
    toggleMissingMandatoryComponentsAlert("hide");
}

function validateInput() {
    let missingMandatoryInputs = getMissingMandatoryInputs();

    if (missingMandatoryInputs.length > 0) {
        processInputError(missingMandatoryInputs);
        return false
    } else {
        processInputPass();
        return true
    }
}
