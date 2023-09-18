package main

import (
	"fmt"

	"github.com/rprata/macgo/generator"
	"github.com/rprata/macgo/vendor"
)

func main() {
	macAddress := generator.GenerateRandomMACAddress()
	fmt.Println(macAddress)
	fmt.Println(vendor.GetVendorInfo(macAddress))

	fmt.Println(vendor.GetVendorInfo("98:CA:33:AA:BB:CC"))
	fmt.Println(vendor.GetPrefixList("raspberry"))

	fmt.Println(generator.GenerateRandomMACAddressByVendor("raspberry"))
}
