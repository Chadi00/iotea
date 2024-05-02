package utils

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type DeviceInfo struct {
	IP   string `json:"ip"`
	MAC  string `json:"mac"`
	Name string `json:"name"`
}

func ScanNetwork() ([]DeviceInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "arp", "-a")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running arp:", err)
		return nil, err
	}

	return parseOutput(out.String())
}

func parseOutput(output string) ([]DeviceInfo, error) {
	lines := strings.Split(output, "\n")
	var devices []DeviceInfo
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 3 {
			ip := strings.Trim(fields[1], "()")
			mac := fields[3]
			if mac != "(incomplete)" {
				name := "-"
				devices = append(devices, DeviceInfo{IP: ip, MAC: mac, Name: name})
			}
		}
	}
	return devices, nil
}
