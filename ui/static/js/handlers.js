function submitButtonHandler() {

    async function createRequestStartFetchingChain() {
        hidePreviousErrors();

        let inputValidationErrors = getInputValidationErrors();
        if (inputValidationErrors.length > 0) {
            DOMUpdateOnInputError(inputValidationErrors);
            return;
        }

        let requestId = null;
        try {
            const objectId = await _fetch_post_new_request();
            requestId = objectId["RequestID"];
        } catch (e) {
            DOMUpdateOnBackendServiceError(JSON.stringify(e));
            return;
        }

        DOMUpdateOnBackendServiceFetchingDataStart();
        try {
            await _update_results_table_with_fetched_data(requestId);
        } catch (e) {
            DOMUpdateOnBackendServiceError(JSON.stringify(e));
            return;
        }
        DOMUpdateOnBackendServiceFetchingDataEnd();

    }

    createRequestStartFetchingChain().catch();
}
