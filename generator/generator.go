package generator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rprata/macgo/lookup"
)

// GenerateRandomMACAddress generates a random MAC address (48 or 64 bits)
func GenerateRandomMACAddress(bits uint8) (string, error) {
	if bits != 48 && bits != 64 {
		return "", fmt.Errorf("ERROR: Invalid number of bits")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var macBytes []byte

	if bits == 48 {
		macBytes = make([]byte, 6)
	} else { // 64 bits
		macBytes = make([]byte, 8)
	}

	for i := 0; i < len(macBytes); i++ {
		macBytes[i] = byte(r.Intn(256))
	}

	// Defina o bit U/L para indicar localmente administrado (em oposição a globalmente único)
	macBytes[0] |= 2

	macAddress := fmt.Sprintf("%02X:%02X:%02X:%02X:%02X:%02X",
		macBytes[0], macBytes[1], macBytes[2], macBytes[3], macBytes[4], macBytes[5])

	if bits == 64 {
		macAddress = fmt.Sprintf("%s:%02X:%02X", macAddress, macBytes[6], macBytes[7])
	}

	return macAddress, nil
}

// GenerateRandomMACAddressByVendor generates a random MAC address for a given vendor name
func GenerateRandomMACAddressByVendor(vendorName string) string {
	if vendorPrefixArr := lookup.GetPrefixList(vendorName); len(vendorPrefixArr) > 0 {
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
