package core

import (
	"os/exec"
	"time"

	"gobox/utils"
)

func RunScript(scriptPath string, listArgs []string) {
	// script := exec.Command(scriptPath, listArgs...)

	// out, err := script.Output()
	// if err != nil {
	// 	fmt.Println("could not run command: ", err)
	// 	return
	// }

	// fmt.Printf("Output:\n%s\n", string(out))

	start := time.Now()
	cmd := exec.Command(scriptPath, listArgs...)
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
		result.ExitCode = -1 // tu peux récupérer le vrai code si besoin
	} else {
		result.ExitCode = 0
	}

	result.PrintPretty()
}
