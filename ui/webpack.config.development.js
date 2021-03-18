const webpack = require("webpack");

module.exports = () => ({
    entry: './javascript_src/index.js',
    output: {
        filename: './bundle.js'
    },
    plugins: [
        new webpack.DefinePlugin({
            "process.env.API_URL": JSON.stringify("http://localhost:8080/api/v1/requests/"),
            "process.env.API_DOC_MD_LOCATION": JSON.stringify("http://localhost:8080/api_documentation.md")
        })

    ]

});
