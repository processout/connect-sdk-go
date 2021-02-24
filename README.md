# Ingenico Connect Go SDK

## Introduction

The Go SDK helps you to communicate with the [Ingenico Connect](https://epayments.developer-ingenico.com/) Server API. It's primary features are:

* convenient go wrapper around the API calls and responses
    * marshalls Go request structs to HTTP requests
    * unmarshalls HTTP responses to Go response structs or Go exceptions
* handling of all the details concerning authentication
* handling of required meta data

Its use is demonstrated by an example for each possible call. The examples execute a call using the provided API keys.

See the [Ingenico Connect Developer Hub](https://epayments.developer-ingenico.com/documentation/sdk/server/go/) for more information on how to use the SDK.

## Requirements

Go version 1.11 or higher is required. No additional packages are required.

## Examples

This repository also contains some example code. This code is contained in the examples folder.

## Installation

### Source

To install the latest version of this repository, run the following command from a terminal:

    go get github.com/processout/connect-sdk-go

### Release

Go 1.11 added [module support](https://blog.golang.org/using-go-modules) and with that support for versions in `go get`. This means that, if your project uses modules, you can add `@version` to the go get command to get a specific version. For example, `go get github.com/processout/connect-sdk-go@2.9.0` will download version 2.9.0 of the SDK. See the releases page for an overview of available releases.

If your project does not use modules yet, you will need to use the instructions above to install from source. Note that new major versions may introduce breaking changes. We therefore recommend using modules in your project. See [Migrating to Go Modules](https://blog.golang.org/migrating-to-go-modules) for more information.

## Running tests

There are two types of tests: unit tests and integration tests. The unit tests will work out-of-the-box; for the integration tests some configuration is required. First, some environment variables need to be set:

* `connect.api.apiKeyId` for the API key id to use. This can be retrieved from the Configuration Center.
* `connect.api.secretApiKey` for the secret API key to use. This can be retrieved from the Configuration Center.
* `connect.api.merchantId` for your merchant ID.
* `connect.api.proxyUrl` for the URL to the proxy to use (optional). If set, it should be in the form `scheme://[userinfo@]host[:port]`. Examples: `http://proxy.example.org`, `http://user:pass@proxy.example.org`, `http://proxy.example.org:3128`.

The following commands can now be executed from the root directory of the SDK folder to execute the tests:

* Unit tests:
    
    ```
    go test ./...
    ```
*  Both unit and integration tests:
    
    ```
    go test -tags=integration  ./...
    ```

## Processout Specific changes

In order to make this repo work for our integration a few changes have been made. These are the following:
* In `factory.go` - modifications were made so that we could create a client by passing in our own HTTP client
* In `factory.go` - further modifications were made so that we can pass in a custom host. Previously this was always defaulting to the production host. Now we can send in a custom host to access the sandbox environment.
* As the response structs use pointers - we needed to have getters to safely access variables. We make use of `gen-accessors.go` within the `domain` folder in order to generate these. If in the future new structs needed to be added within the domain folder - you can simply add the following to the top of the file: `//go:generate go run ../gen-accessors.go` and run `go generate ./...` - this will auto generate the accessors for you to use. This file was forked from: https://github.com/google/go-github/blob/master/github/gen-accessors.go (We've turned off generate the automated tests for these)