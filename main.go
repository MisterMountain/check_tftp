package main

import (
	"fmt"
	"os"

	"github.com/NETWAYS/go-check"
	"github.com/pin/tftp/v3"
	log "github.com/sirupsen/logrus"
)

func main() {
	defer check.CatchPanic()
	config := check.NewConfig()
	config.Name = "check_test"
	config.Readme = `Test Plugin`
	config.Version = "0.0.1"
	config.Timeout = 10

	value := config.FlagSet.IntP("value", "", 10, "test value")
	warning := config.FlagSet.IntP("warning", "w", 20, "warning threshold")
	critical := config.FlagSet.IntP("critical", "c", 50, "critical threshold")
	hostname := config.FlagSet.StringP("hostname", "H", "", "hostname of TFTP Server")
	file := config.FlagSet.StringP("file", "f", "", "file to receive")

	config.ParseArguments()

	log.Info("Start logging")

	// print help if none flag is set
	if config.FlagSet.NFlag() == 0 {
		os.Exit(0)
	}

	// Parse the TFTP server address from the hostname flag
	address := fmt.Sprintf("%s:69", *hostname)

	// Open a connection to the TFTP server
	client, err := tftp.NewClient(address)
	if err != nil {
		check.Exitf(check.Critical, "Failed to create TFTP client: %v", err)
	}

	// Download the file
	wt, err := client.Receive(*file, "octet")
	if err != nil {
		check.Exitf(check.Critical, "Failed to start TFTP transfer: %v", err)
	}

	f, err := os.Create(*file)
	if err != nil {
		check.Exitf(check.Critical, "Failed to create file: %v", err)
	}
	defer f.Close()

	n, err := wt.WriteTo(f)
	if err != nil {
		check.Exitf(check.Critical, "Failed to download file: %v", err)
	}

	check.Exitf(check.OK, "%d bytes received", n)

	// time.Sleep(20 * time.Second)

	if *value > *critical {
		check.Exitf(check.Critical, "value is %d", *value)
	} else if *value > *warning {
		check.Exitf(check.Warning, "value is %d", *value)
	} else {
		check.Exitf(check.OK, "value is %d", *value)
	}
}
