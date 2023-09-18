# MACGo: A Go Package for Handling MAC Addresses

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

MACGo is a Go package that provides tools for working with MAC addresses. It simplifies common tasks such as generating random MAC addresses, validating MAC addresses, and retrieving the vendor information associated with a MAC address.

## Features

- Generate random MAC addresses.
- Check the validity of a MAC address.
- Retrieve the vendor information for a MAC address.

## Installation

To install the MACGo package, use the `go get` command:

```bash
go get github.com/rprata/macgo
```

## Usage

### Generate a Random MAC Address

```go
package main

import (
    "fmt"
    "github.com/rprata/macgo/generator"
)

func main() {
    randomMAC := macgo.GenerateRandomMAC()
    fmt.Println("Random MAC Address:", randomMAC)
    randomRaspberryPiMAC := generator.GenerateRandomMACAddressByVendor("raspberry")
    fmt.Println("Random MAC Address (Raspberry PI): ", randomRaspberryPiMAC)

}
```

### Check if a MAC Address is Valid

```go
package main

import (
    "fmt"
    "github.com/rprata/macgo/checker"
)

func main() {
    macAddress := "00:1A:2B:3C:4D:5E"
    
    if checker.IsValidMACAddress(macAddress) {
        fmt.Printf("%s is a valid MAC address\n", macAddress)
    } else {
        fmt.Printf("%s is not a valid MAC address\n", macAddress)
    }
}
```

### Get Vendor Information by MAC Address

```go
package main

import (
    "fmt"
    "github.com/rprata/macgo/lookup"
)

func main() {
    macAddress := "00:1A:2B:3C:4D:5E"
    vendorInfo, err := macgo.GetVendorInfo(macAddress)
    
    if err != nil {
        fmt.Printf("Error getting vendor info: %s\n", err)
    } else {
        fmt.Printf("Vendor: %s\n", vendorInfo)
    }
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions to macgo are welcome! Please read our [Contribution Guidelines](CONTRIBUTING.md) for more information on how to get involved.
```

You can copy the above Markdown content and paste it into your project's README file or Markdown editor.