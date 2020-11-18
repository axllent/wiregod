package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// WGStatus return the current status
func WGStatus() error {
	connections := WGConnections()

	if len(connections) == 0 {
		return fmt.Errorf("WireGuard is not active")
	}

	for _, wg := range connections {
		output, err := SudoExec("wg show " + wg)
		if err != nil {
			return err
		}
		fmt.Printf("WireGuard interface \"%s\" up:\n\n%s\n", wg, output)
	}

	ip, err := PublicIP()
	if err != nil {
		return fmt.Errorf("Error fetching public IP: %s", err)
	}

	fmt.Printf("Public IP: %s\n", ip)

	return nil
}

// PublicIP returns your public IP
func PublicIP() (string, error) {

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("https://ipinfo.io/ip")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(body)), nil
}
