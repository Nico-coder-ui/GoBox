package cmd

import (
	"fmt"
	"gobox/core"
	"gobox/utils"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gobox",
	Short: "Binary & script analyser",
	Long: `Binary & script analyser.
Example :
gobox ./my_script.sh --args="val1 val2"`,
	RunE: goboxHandle,
}

func goboxHandle(cmd *cobra.Command, args []string) error {
	err := core.DockerSandboxImage("gobox_sandbox:1.0", "/home/asure/Personal/Code/GoBox/Dockerfile")
	if err != nil {
		return err
	}

	if len(args) == 0 {
		return fmt.Errorf("you have to right a script filepath first")
	}

	scriptFilepath := args[0]
	if !utils.IsLocally(scriptFilepath) {
		return fmt.Errorf("your script has to be accessible")
	}
	longArgs, _ := cmd.Flags().GetString("args")
	newArgs := strings.Fields(longArgs)

	fmt.Printf("Filepath: %s\nArgs: %s\n\n\n", scriptFilepath, newArgs)
	core.RunScript(scriptFilepath, newArgs)
	return nil
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("args", "a", "", `Arguments of the script (ex : "val1 val2")`)
}
