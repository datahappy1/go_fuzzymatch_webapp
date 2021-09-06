const webpack = require("webpack");

module.exports = () => ({
    entry: './src/index.js',
    output: {
        filename: './bundle.js'
    },
    plugins: [
        new webpack.DefinePlugin({
            "process.env.API_URL": JSON.stringify("https://fuzzster.herokuapp.com/api/v1/requests/"),
            "process.env.API_DOC_MD_LOCATION": JSON.stringify("https://fuzzster.herokuapp.com/api_documentation.md")
        })

    ]

});
