package handlers

import (
	"os"
)

// GetPeerList Obtains Peer List
// From Environment Variable
func GetPeerList() []string {
	// if os.Getenv("PEERS") == "" {
	// 	return []string{}
	// }
	return []string{"peer-1"}
}

// GetNetwork Obtains Network
// From Environment Variable
func GetNetwork() string {
	return os.Getenv("NETWORK") + ":8080"
}
