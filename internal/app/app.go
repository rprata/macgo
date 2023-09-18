package app

import (
	"fmt"

	"github.com/rprata/macgo/generator"
	"github.com/urfave/cli"
)

// Generate returns a new cli.App with our desired configuration
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "MACGO CLI"
	app.Usage = "MAC Address Generator and Vendor Lookup"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "Generate a random MAC address",
			Action:  generateMacAddress,
		},
		{
			Name:    "lookup",
			Aliases: []string{"l"},
			Usage:   "Lookup vendor information for a given MAC address",
			Action:  lookupVendor,
		},
	}
	return app
}

func generateMacAddress(c *cli.Context) {
	macAddress := generator.GenerateRandomMACAddress()
	fmt.Println(macAddress)
}

func lookupVendor(c *cli.Context) {
	// macAddress := c.Args().First()
	// if macAddress == "" {
	// 	fmt.Println("ERROR: Please provide a MAC address")
	// 	return
	// }
	// if !checker.IsValidMACAddress(macAddress) {
	// 	fmt.Println("ERROR: Invalid MAC address")
	// 	return
	// }
	// vendorInfo, err := vendor.GetVendorInfo(macAddress)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(vendorInfo)
}
