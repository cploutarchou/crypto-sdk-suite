# Crypto SDK Suite
[![Go Reference](https://pkg.go.dev/badge/github.com/cploutarchou/crypto-sdk-suite.svg)](https://pkg.go.dev/github.com/cploutarchou/crypto-sdk-suite)

**Crypto SDK Suite** is a comprehensive library developed in Go that provides easy-to-use APIs for various cryptocurrency exchanges and services, including Binance, Bybit, CoinMarketCap, and more. This suite is designed to simplify the integration and interaction with different crypto platforms, offering unified access to trading, market data, and account management functionalities. It's perfect for developers building trading bots, market analysis tools, or any application that requires seamless access to multiple crypto services.

### Features

- **Unified API Access**: Connect to multiple cryptocurrency exchanges and services through a single, unified API interface.
- **Trading Operations**: Execute trades, manage orders, and monitor your trading activity across different platforms.
- **Market Data**: Access real-time and historical market data for comprehensive analysis and strategy development.
- **Account Management**: Manage your account details, balances, and transaction history seamlessly.

### Supported Platforms

- Binance
- Bybit
- CoinMarketCap
- And more to come!

### Getting Started

To get started with Crypto SDK Suite, check out our [documentation](https://pkg.go.dev/github.com/cploutarchou/crypto-sdk-suite) for detailed guides and API references.

### Installation

To install the library, run:

```bash
go get github.com/cploutarchou/crypto-sdk-suite
```

### Usage
Here is a basic example to get you started:

```go
package main

import (
    "fmt"
    "github.com/cploutarchou/crypto-sdk-suite"
)

func main() {
    client, err := crypto.NewPublicClient(false, "binance")
    if err != nil {
        log.Fatalf("Failed to initialize client: %v", err)
    }
    
    // Example usage
    data, err := client.GetMarketData("BTCUSDT")
    if err != nil {
        log.Fatalf("Error fetching market data: %v", err)
    }

    fmt.Println(data)
}

```

**Note**: This project is a work in progress. We are continuously adding new features and improving the existing ones to make developers' lives easier.

**Contributions are welcome!** If you'd like to contribute, please feel free to fork the repository and submit pull requests. Your contributions can include adding new features, fixing bugs, or improving the documentation. We appreciate all contributions that help enhance the library's functionality and usability.