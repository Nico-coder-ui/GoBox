package core

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"gobox/utils"
)

func RunScript(scriptPath string, listArgs []string) {
	start := time.Now()

	absPath, err := filepath.Abs(scriptPath)
	if err != nil {
		fmt.Println("Erreur chemin absolu :", err)
		return
	}

	containerScriptPath := "/sandbox/script"

	args := []string{
		"run", "--rm",
		"--network=none",
		"--user=nobody",
		"-v", absPath + ":" + containerScriptPath + ":ro",
		"gobox_sandbox:1.0",
		"/bin/bash", containerScriptPath,
	}
	args = append(args, listArgs...)

	cmd := exec.Command("docker", args...)
	out, err := cmd.CombinedOutput()
	end := time.Now()

	result := utils.ExecutionResult{
		ScriptPath:    scriptPath,
		Args:          listArgs,
		Output:        string(out),
		StartedAt:     start,
		FinishedAt:    end,
		ExecutionTime: end.Sub(start),
		Success:       err == nil,
	}

	if err != nil {
		result.ErrorMsg = err.Error()
		result.ExitCode = -1
	} else {
		result.ExitCode = 0
	}

	result.PrintPretty()
}
