package server

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const AppName = "TaskWeaver"

var (
	// subsequently get from branch or tie to dev & main branch
	Version = "dev"
)

func init() {
	// Get the version from git tags
	if Version == "dev" {
		version, err := getVersionFromGit()
		if err != nil {
			log.Printf("Failed to get version from git")
			// log.Panic("Error getting version from git:", err)
			// don't panic - fallback on dev
		}
		Version = version
	}
}

// getVersionFromGit retrieves the version from the latest git tag
func getVersionFromGit() (string, error) {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Trim leading/trailing whitespaces and newline characters
	version := strings.TrimSpace(string(output))
	return version, nil
}

func FullVersion() string {
	return fmt.Sprintf("%s version %s", AppName, Version)
}
