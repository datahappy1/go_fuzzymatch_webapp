const mandatoryInputIDs = ['stringsToMatch', 'stringsToMatchIn'];

function isValidMandatoryComponent(element) {
    let validatedComponent = document.getElementById(element);

    return !(validatedComponent.value.length === 0);
}

function validateMandatoryInputs() {
    let validationResult = "";

    for (let i = 0; i < mandatoryInputIDs.length; i++) {
        if (isValidMandatoryComponent(mandatoryInputIDs[i]) === false) {
            validationResult += mandatoryInputIDs[i]
        }
    }

    return validationResult;
}
