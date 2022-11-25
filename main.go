package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "goreleaser-testing",
	Long: "Root command",
}

var HelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "hello",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello!!!")
	},
}

var BeforeCmd = &cobra.Command{
	Use:   "before",
	Short: "before",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("before!!!")
	},
}

var ByeCmd = &cobra.Command{
	Use:   "bye",
	Short: "azioncli goodbye",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bye!!!")
	},
}

var CompletionCmd = &cobra.Command{
	Use:                   "completion [bash|zsh|fish|powershell]",
	Short:                 "Generate completion script",
	Long:                  "To load completions",
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			err := cmd.Root().GenBashCompletion(os.Stdout)
			if err != nil {
				return
			}
		case "zsh":
			err := cmd.Root().GenZshCompletion(os.Stdout)
			if err != nil {
				return
			}
		case "fish":
			err := cmd.Root().GenFishCompletion(os.Stdout, true)
			if err != nil {
				return
			}
		case "powershell":
			err := cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
			if err != nil {
				return
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(HelloCmd, CompletionCmd, BeforeCmd)
	HelloCmd.AddCommand(ByeCmd)
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
