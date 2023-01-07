package main

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var waitFlag bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gork [-w] -- command args...",
	Short: "`fork` written in go",
	Long: `gork spawns and disowns any process passed to it as subcommand.
It frees the user from unnecessary hassle like "&!" in linux shells.

For example this:
$ gork google-chrome github.com

directly launches chrome and opens github in a new window`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		subcmd := exec.Command(args[0], args[1:]...)
		var err error
		err = subcmd.Start()
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
		if waitFlag {
			err = subcmd.Wait()
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
		}
	},
	DisableFlagsInUseLine: true,
	SilenceUsage:          true,
	Version:               "1.0",
}

// add all child commands to the root command and sets flags appropriately.
// It only needs to happen once to the rootCmd.
func main() {
	cobra.MinimumNArgs(1)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&waitFlag, "wait", "w", false, "Wait until <command> completes")
}
