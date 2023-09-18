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

## MACGo CLI

The MACGo CLI is a command-line tool designed to provide functionalities related to MAC addresses, such as generation and vendor information lookup. It leverages the following Go packages for its operations:

- **macgo/checker**: A package for validating MAC addresses.
- **macgo/generator**: A package for generating random MAC addresses.
- **macgo/lookup**: A package for retrieving vendor information by MAC address.

### Usage

The MACGo CLI offers the following commands:

#### 1. Generate a Random MAC Address

Use the `generate` command to generate a random MAC address. You can also specify a vendor name to generate a MAC address associated with that vendor.

```sh
macgo generate [-v, --vendor <vendor-name>]
```

- `-v, --vendor <vendor-name>`: Optional flag to specify a vendor name for generating a MAC address.

Example:

```sh
macgo generate
macgo generate -v "Cisco"
```

#### 2. Lookup Vendor Information

Use the `lookup` command to retrieve vendor information for a given MAC address.

```sh
macgo lookup <mac-address>
```

- `<mac-address>`: The MAC address for which you want to retrieve vendor information.

Example:

```sh
macgo lookup 00:1A:2B:3C:4D:5E
```

### Version

The current version of the MACGo CLI is "0.0.1."

### How to Run

To run the MACGo CLI, build and execute the binary. Here are the general steps:

1. Install the CLI:
   ```sh
   go install yourproject/app
   ```

2. Run the CLI:
   ```sh
   macgo <command>
   ```

Replace `<command>` with one of the available commands, such as `generate` or `lookup`.

### Examples

Here are some examples of how to use the MACGo CLI:

- Generate a random MAC address:
  ```sh
  macgo generate
  ```

- Generate a random MAC address for a specific vendor (e.g., Cisco):
  ```sh
  macgo generate -v "Cisco"
  ```

- Lookup vendor information for a MAC address (e.g., 00:1A:2B:3C:4D:5E):
  ```sh
  macgo lookup 00:1A:2B:3C:4D:5E
  ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions to macgo are welcome! Please read our [Contribution Guidelines](CONTRIBUTING.md) for more information on how to get involved.
```

You can copy the above Markdown content and paste it into your project's README file or Markdown editor.