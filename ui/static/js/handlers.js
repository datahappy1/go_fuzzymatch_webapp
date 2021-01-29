function submitButtonHandler() {

    if (!validateInput()) {
        return
    }

    clearResultsTable();

    createRequestStartFetchingChain();

    showResults();

    jumpToAnchor("results");

}
