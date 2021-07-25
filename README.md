# 0xScan
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/0xTrackerApp/0xScan/blob/master/LICENSE)
[![testing](https://github.com/0xTrackerApp/0xScan/actions/workflows/testing.yml/badge.svg)](https://github.com/0xTrackerApp/0xScan/actions/workflows/testing.yml)

## 💡 About
This module offers an API wrapper for he blockexplorers from the etherscan.io team.

The APIs from the following explorers are supported at the moment:
- https://etherscan.io
- https://bscscan.com
- https://polygonscan.com
- https://ftmscan.com

## 🚀 Getting started

### 🧑‍💻 Create a client

```go
package main

import (
    "context"
    xscan "github.com/0xTrackerApp/0xScan"
)

func main() {
    // create empty context interface
    ctx := context.Background()

    // Create a new client
    client := xscan.NewClient(xscan.BscOpts, "<yourApiKey>") // other available options are xscan.EthOpts, xscan.PolygonOpts, etc

    _ = client
}
```

## 📜 Licensing
This SDK is released under the MIT-License found in the [LICENSE](https://github.com/0xTrackerApp/0xScan/blob/master/LICENSE) file.