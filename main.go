package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"

	"github.com/NETWAYS/go-check"
	"github.com/pin/tftp/v3"
)

func main() {
	defer check.CatchPanic()
	config := check.NewConfig()
	config.Name = "check_tftp"
	config.Readme = `TFTP Check Plugin`
	config.Version = "0.0.1"
	config.Timeout = 10

	hostname := config.FlagSet.StringP("hostname", "H", "", "hostname of TFTP Server")
	file := config.FlagSet.StringP("file", "f", "", "file to receive")
	checksum := config.FlagSet.StringP("checksum", "C", "", "SHA1 checksum of file")

	config.ParseArguments()

	if *hostname == "" || *file == "" {
		var missing string
		if *hostname == "" && *file == "" {
			missing = "hostname and file"
		}
		if *hostname == "" && *file != "" {
			missing = "hostname"
		}
		if *hostname != "" && *file == "" {
			missing = "file"
		}
		config.FlagSet.Usage()
		check.Exitf(check.Unknown, "Missing required arguments: %s", missing)
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

	if *checksum != "" {
		verifyChecksum(*file, *checksum)
	}

	check.Exitf(check.OK, "%d bytes received", n)
}

func verifyChecksum(file string, checksum string) bool {
	f, err := os.Open(file)
	if err != nil {
		check.Exitf(check.Critical, "Failed to open file: %v", err)
	}
	defer f.Close()

	h := sha1.New()
	if _, err := f.Seek(0, 0); err != nil {
		check.Exitf(check.Critical, "Failed to seek file: %v", err)
	}
	if _, err := io.Copy(h, f); err != nil {
		check.Exitf(check.Critical, "Failed to checksum file: %v", err)
	}
	if fmt.Sprintf("%x", h.Sum(nil)) != checksum {
		check.Exitf(check.Critical, "SHA1 hash mismatch: expected %s, got %s", checksum, fmt.Sprintf("%x", h.Sum(nil)))
	}
	return true
}
