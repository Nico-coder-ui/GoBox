package core

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func DockerSandboxImage(tag string, dockerfilePath string) error {
	checkCmd := exec.Command("docker", "images", "-q", tag)
	var out bytes.Buffer
	checkCmd.Stdout = &out
	err := checkCmd.Run()
	if err != nil {
		return fmt.Errorf("failed to check image: %w", err)
	}

	if strings.TrimSpace(out.String()) != "" {
		return nil
	}

	buildCmd := exec.Command(
		"docker", "build", "-t", tag, "-f", dockerfilePath, ".",
	)
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr

	err = buildCmd.Run()
	if err != nil {
		return fmt.Errorf("failed to build image: %w", err)
	}

	return nil
}
