package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var waitFlag bool
var curDir string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gork [-w] [-C cwd] command args...",
	Short: "`fork` written in go",
	Long: `gork spawns and disowns any process passed to it as subcommand.
It frees the user from unnecessary hassle like "&!" in linux shells.

For example this:
$ gork google-chrome github.com

directly launches chrome and opens github in a new window`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		if curDir != "." {
			err = os.Chdir(curDir)
			if err != nil {
				return err
			}
		}
		subcmd := exec.Command(args[0], args[1:]...)
		err = subcmd.Start()
		if err != nil {
			return err
		}
		if waitFlag {
			err = subcmd.Wait()
			if err != nil {
				return err
			}
		}
		return nil
	},
	DisableFlagsInUseLine: true,
	SilenceUsage:          true,
	Version:               "1.4",
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
	rootCmd.Flags().StringVarP(&curDir, "dir", "C", ".", "start <command> from `dir` instead pf current directory")
	rootCmd.Flags().SetInterspersed(false)
	rootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		return fmt.Errorf("%w\nUsage: %s", err, cmd.UseLine())
	})
}
