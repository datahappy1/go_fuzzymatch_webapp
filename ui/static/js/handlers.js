function submitButtonHandler() {
    if (isValidMandatoryComponents() === false) {
        toggleMissingMandatoryComponentsAlert("show");
        return
    }
    toggleMissingMandatoryComponentsAlert("hide");

    clearResultsTable();
    createRequestStartFetchingChain();
    showResults();
}
