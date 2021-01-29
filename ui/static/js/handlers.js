function submitButtonHandler() {

    let inputValidationErrors = getInputValidationErrors();

    if (inputValidationErrors.length > 0) {
        processInputError(inputValidationErrors);
        return
    }
    processInputPass();

    clearResultsTable();

    createRequestStartFetchingChain();

    showResults();

    jumpToAnchor("results");

}
