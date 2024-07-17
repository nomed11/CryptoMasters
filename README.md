# CryptoMasters API
=================

Welcome to the **CryptoMasters API**! This project fetches real-time cryptocurrency exchange rates from CEX.io using Go, and was made to understand concurrent programming and robust error handling in Go.

## Table of Contents
-----------------

-   [Features](#features)
-   [Installation](#installation)
-   [Usage](#usage)
-   [Project Structure](#project-structure)
-   [CI/CD with GitHub Actions](#cicd-with-github-actions)
-   [License](#license)

## Features

-   **Concurrent API Calls**: Fetches data concurrently using goroutines and `sync.WaitGroup`.
-   **Robust Error Handling**: Handles network issues, invalid responses, and empty currency inputs gracefully.
-   **JSON Parsing**: Efficiently parses JSON responses from the CEX.io API.
-   **Automated CI/CD**: Continuous integration and testing with GitHub Actions.

## Installation

1\. **Clone the repository**:\
   ```sh\
   git clone https://github.com/yourusername/cryptomasters-api.git\
   cd cryptomasters-api\
   ```\
2\. **Install dependencies**:\
   ```sh\
   go mod tidy\
   ```

## Usage

1\. **Run the application**:\
   ```sh\
   go run main.go\
   ```\
   Example output:\
   ```\
   Currency: BTC, Price: $XXXXX.XX\
   Currency: ETH, Price: $XXX.XX\
   Currency: XRP, Price: $X.XX\
   ```

2\. **Run tests**:\
   ```sh\
   go test ./...\
   ```

## Project Structure

```\
cryptomasters-api/\
├── api/\
│   ├── cex.go\
│   ├── responses.go\
│   └── cex_test.go\
├── datatypes/\
│   └── data.go\
├── main.go\
├── .github/\
│   └── workflows/\
│       └── go.yml\
├── go.mod\
└── .gitignore\
```

## CI/CD with GitHub Actions

Automated builds and tests are triggered on every push using GitHub Actions.

### GitHub Actions Workflow

```yaml\
name: Go

on:\
  push:\
    branches: [ "main" ]\
  pull_request:\
    branches: [ "main" ]

jobs:\
  build:\
    runs-on: ubuntu-latest\
    steps:\
    - uses: actions/checkout@v4\
    - name: Set up Go\
      uses: actions/setup-go@v4\
      with:\
        go-version: '1.22'\
    - name: Build\
      run: go build -v ./...\
    - name: Test\
      run: go test -v ./...\
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.