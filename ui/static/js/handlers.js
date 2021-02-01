function submitButtonHandler() {

    async function createRequestStartFetchingChain() {
        let inputValidationErrors = getInputValidationErrors();

        if (inputValidationErrors.length > 0) {
            processInputError(inputValidationErrors);
            return
        }
        processInputPass();

        let requestId = null;
        try {
            const ObjectId = await _fetch_post_new_request();
            requestId = ObjectId.get("RequestID");
        } catch (e) {
            processBackendServiceError(e.message);
            return;
        }
        processBackendServicePass();

        processBackendServiceFetchingDataStart();
        try {
            await _update_results_table_with_fetched_data(requestId);
        } catch (e) {
            processBackendServiceError(e);
            return;
        }
        processBackendServicePass();
        processBackendServiceFetchingDataEnd();

    }

    createRequestStartFetchingChain().catch();
}
