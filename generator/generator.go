package generator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rprata/macgo/vendor"
)

// GenerateRandomMACAddress generates a random MAC address
func GenerateRandomMACAddress() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	macBytes := make([]byte, 6)
	for i := 0; i < 6; i++ {
		macBytes[i] = byte(r.Intn(256))
	}

	macBytes[0] &^= 2

	macAddress := fmt.Sprintf("%02X:%02X:%02X:%02X:%02X:%02X",
		macBytes[0], macBytes[1], macBytes[2], macBytes[3], macBytes[4], macBytes[5])

	return macAddress
}

// GenerateRandomMACAddressByVendor generates a random MAC address for a given vendor name
func GenerateRandomMACAddressByVendor(vendorName string) string {
	if vendorPrefixArr := vendor.GetPrefixList(vendorName); len(vendorPrefixArr) > 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		var vendorPrefix string = ""
		for _, item := range vendorPrefixArr {
			if len(item) == 8 {
				vendorPrefix = item
				break
			}
		}

		macBytes := make([]byte, 6)
		for i := 0; i < 3; i++ {
			macBytes[i] = byte(r.Intn(256))
		}

		macBytes[0] &^= 2

		macAddress := fmt.Sprintf("%s:%02X:%02X:%02X",
			vendorPrefix, macBytes[0], macBytes[1], macBytes[2])

		return macAddress
	}

	return ""
}
