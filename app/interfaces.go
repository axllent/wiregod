package app

import (
	"fmt"
	"os"
	"strings"
)

// Configs will return a slice of available wireguard interface configs
func Configs() []string {
	interfaces := []string{}

	files, err := SudoExec("ls /etc/wireguard")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(2)
	}

	configs := strings.Fields(files)
	for _, iface := range configs {
		if strings.HasSuffix(iface, ".conf") {
			p := strings.Split(iface, ".conf")
			interfaces = append(interfaces, p[0])
		}
	}

	return interfaces
}

// WGConnections return running connection interfaces
func WGConnections() []string {
	ifs, err := SudoExec("wg show interfaces")
	if err != nil {
		return []string{}
	}

	return strings.Fields(ifs)
}

// PublicKey returns a public key
func PublicKey(wg string) string {
	out, err := SudoExec("wg show " + wg + " public-key")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return out
}
