package lookup

import (
	"fmt"
	"strings"

	"github.com/rprata/macgo/internal/helper"
)

// GetVendorInfo returns the vendor name for a given MAC address
func GetVendorInfo(macAddress string) (string, error) {
	macs, err := helper.LoadData()
	if err != nil {
		return "", fmt.Errorf("ERROR loading MACs: %s", err)
	}

	for _, item := range macs {
		if strings.HasPrefix(macAddress, item.MacPrefix) {
			return item.VendorName, nil
		}
	}
	return "", fmt.Errorf("ERROR: MAC prefix not found")
}

// GetPrefixList returns the MAC prefix for a given vendor name
func GetPrefixList(vendor string) []string {
	macs, err := helper.LoadData()
	if err != nil {
		panic(fmt.Sprintf("ERROR loading MACs: %s", err))
	}

	var res []string

	for _, item := range macs {
		if strings.Contains(strings.ToUpper(item.VendorName), strings.ToUpper(vendor)) {
			res = append(res, item.MacPrefix)
		}
	}
	return res
}
