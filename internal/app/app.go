package app

import (
	"fmt"

	"github.com/rprata/macgo/checker"
	"github.com/rprata/macgo/generator"
	"github.com/rprata/macgo/lookup"
	"github.com/urfave/cli"
)

// Generate returns a new cli.App with our desired configuration
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "MACGo CLI"
	app.Usage = "MAC Address Generator and Vendor Lookup"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "Generate a random MAC address",
			Action:  generateMacAddress,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "vendor, v",
					Usage: "Generate a random MAC address for a given vendor name",
				},
			},
		},
		{
			Name:    "lookup",
			Aliases: []string{"l"},
			Usage:   "Lookup vendor information for a given MAC address",
			Action:  lookupVendor,
		},
		{
			Name:    "check",
			Aliases: []string{"c"},
			Usage:   "Check if a given MAC address is valid",
			Action:  isValidMACAddress,
		},
	}
	return app
}

func generateMacAddress(c *cli.Context) {
	vendorName := c.String("vendor")
	var macAddress string = ""
	if vendorName != "" {
		macAddress = generator.GenerateRandomMACAddressByVendor(vendorName)
	} else {
		macAddress = generator.GenerateRandomMACAddress()
	}
	fmt.Println(macAddress)
}

func lookupVendor(c *cli.Context) {
	macAddress := c.Args().First()
	if macAddress == "" {
		fmt.Println("ERROR: Please provide a MAC address")
		return
	}
	if !checker.IsValidMACAddress(macAddress) {
		fmt.Println("ERROR: Invalid MAC address")
		return
	}
	vendorInfo, err := lookup.GetVendorInfo(macAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(vendorInfo)
}

func isValidMACAddress(c *cli.Context) {
	macAddress := c.Args().First()
	if macAddress == "" {
		fmt.Println("ERROR: Please provide a MAC address")
		return
	}
	if checker.IsValidMACAddress(macAddress) {
		fmt.Println("Valid MAC address")
	} else {
		fmt.Println("Invalid MAC address")
	}
}
