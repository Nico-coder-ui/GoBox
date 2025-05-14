package cmd

import (
	"strings"

	"gobox/core"

	"github.com/spf13/cobra"
)

var runAscript = &cobra.Command{
	Use:   "run",
	Short: "Run a scipt",
	Long: `Run a script with gobox and will be get some logs
For exemple :

gobox run [script] [args...]`,
	Run: runScript,
}

func init() {
	rootCmd.AddCommand(runAscript)

	runAscript.Flags().StringP("path", "p", "", "Script's path")
	runAscript.Flags().StringP("args", "a", "", "Args of the prog (ex : \"arg1 arg2 arg3\")")
}

func runScript(cmd *cobra.Command, args []string) {
	scriptPath, _ := cmd.Flags().GetString("path")

	scriptArgs, _ := cmd.Flags().GetString("args")
	listArgs := strings.Fields(scriptArgs)

	core.RunScript(scriptPath, listArgs)
}
