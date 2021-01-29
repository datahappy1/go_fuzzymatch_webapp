function submitButtonHandler() {

    let validatedMandatoryInputs = validateMandatoryInputs();
    if (validatedMandatoryInputs.length > 0) {
        updateMissingMandatoryComponentsAlert(validatedMandatoryInputs);
        toggleMissingMandatoryComponentsAlert("show");
        return
    }
    toggleMissingMandatoryComponentsAlert("hide");

    clearResultsTable();

    createRequestStartFetchingChain();

    showResults();

}
