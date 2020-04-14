package main

import (
	"strings"
	"testing"
)

func TestNetworkInfoCommand(t *testing.T) {
	networkInfo := new(NetworkInfo)
	networkInfo.Command = "GUIDER test"
	cmd := getRunGuiderCommand(networkInfo)

	if !strings.Contains(cmd, "run") {
		t.Error("cmd not contain run")
	}
}