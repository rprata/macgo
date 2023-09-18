package checker

import "regexp"

// Checker is a function that checks if a MAC address is valid
func IsValidMACAddress(macAddress string) bool {
	macPattern := `^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$|^([0-9A-Fa-f]{2}:){7}([0-9A-Fa-f]{2})$`
	re := regexp.MustCompile(macPattern)
	return re.MatchString(macAddress)
}
