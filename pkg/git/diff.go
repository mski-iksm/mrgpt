package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GitDiff() string {
	cmd := exec.Command("git", "diff", "origin/master..HEAD")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Stdout: %s\n", stdout.String())
		fmt.Printf("Stderr: %s\n", stderr.String())
	}
	return stdout.String()
}
