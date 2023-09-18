package helper

import (
	"encoding/json"
	"fmt"

	"github.com/rprata/macgo/internal/data"
)

type MACInfo struct {
	MacPrefix  string `json:"macPrefix"`
	VendorName string `json:"vendorName"`
	Private    bool   `json:"private"`
	BlockType  string `json:"blockType"`
	LastUpdate string `json:"lastUpdate"`
}

func LoadData() ([]MACInfo, error) {
	var macInfo []MACInfo
	if err := json.Unmarshal([]byte(data.GetMACData()), &macInfo); err != nil {
		return nil, fmt.Errorf("ERROR unmarshaling JSON: %s", err)
	}

	return macInfo, nil
}
