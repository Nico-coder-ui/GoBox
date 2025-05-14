package utils

import (
	"fmt"
	"time"
)

type ExecutionResult struct {
	ScriptPath    string        `json:"script_path"`
	Args          []string      `json:"args"`
	ExitCode      int           `json:"exit_code"`
	Output        string        `json:"output"`
	ErrorMsg      string        `json:"error,omitempty"`
	Success       bool          `json:"success"`
	StartedAt     time.Time     `json:"started_at"`
	FinishedAt    time.Time     `json:"finished_at"`
	ExecutionTime time.Duration `json:"execution_time"`
}

func (r *ExecutionResult) PrintPretty() {
	fmt.Println("📄 Script exécuté :", r.ScriptPath)
	fmt.Println("🧩 Arguments      :", r.Args)
	fmt.Println("🕓 Début          :", r.StartedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("🕔 Fin            :", r.FinishedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("⏱️ Temps total    :", r.ExecutionTime)

	if r.Success {
		fmt.Println("✅ Succès         : Oui")
	} else {
		fmt.Println("❌ Succès         : Non")
		fmt.Println("🚨 Erreur         :", r.ErrorMsg)
	}

	fmt.Println("💬 Sortie :")
	fmt.Println("-------------------")
	fmt.Println(r.Output)
	fmt.Println("-------------------")
}
