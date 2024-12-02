
# Transaction Parser Service

## Description

This service is designed to parse transactions from the Ethereum blockchain. It identifies whether a subscribed address, specified by the client, is involved in a transaction. If so, the transaction data is stored in the service.

### Key Features:
1. Parses blocks from the Ethereum blockchain.
2. Checks transactions for the involvement of subscribed addresses.
3. Stores transaction data locally if a match is found.

## Getting Started

### How to Run

1. **Ensure Configuration**:
    - The application uses a configuration file that can be specified via flags. Review and update the configuration settings in the `config` directory.
    - Supported configuration files include:
        - `config.go`: Contains default configuration struct definitions.
        - `yaml`: The YAML configuration file with runtime parameters.

2. **Build and Run**:
    - Navigate to the project directory containing the `cmd/transaction-parser-api` folder.
    - Use the Go compiler to build and run the application:
      ```bash
      go run ./cmd/transaction-parser-api/main.go --config <path_to_config.yaml>
      ```
    - Replace `<path_to_config.yaml>` with the actual path to your YAML configuration file.

3. **Environment Variables**:
    - Ensure that the `EthereumUrl` is correctly set in the configuration to connect to the Ethereum node.

4. **Dependencies**:
    - This project does not require any external Go module dependencies beyond those specified in `go.mod`.

## Contributing

## Authors

## License

Feel free to reach out with questions or suggestions for improvement.
