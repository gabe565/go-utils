package cobrax

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var ErrInvalidShell = errors.New("invalid shell")

type Shell string

const (
	FlagCompletion = "completion"

	Bash       Shell = "bash"
	Zsh        Shell = "zsh"
	Fish       Shell = "fish"
	PowerShell Shell = "powershell"
)

// RegisterCompletionFlag registers a flag to generate a shell completion script.
// Cobra typically uses a subcommand for this, but this flag is useful for
// executables that don't have any subcommands.
//
// An initializer is registered which will trigger when the flag is set.
func RegisterCompletionFlag(cmd *cobra.Command) error {
	cmd.Flags().String(FlagCompletion, "", "Generate the autocompletion script for the specified shell (one of bash, zsh, fish, powershell)")

	if err := cmd.RegisterFlagCompletionFunc(FlagCompletion, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return []string{string(Bash), string(Zsh), string(Fish), string(PowerShell)}, cobra.ShellCompDirectiveNoFileComp
	}); err != nil {
		return err
	}

	cobra.OnInitialize(func() {
		f := cmd.Flags().Lookup(FlagCompletion)
		if f == nil || f.Value.String() == "" {
			return
		}

		if err := GenCompletion(cmd, Shell(f.Value.String())); err != nil {
			// Cobra initializers can't return errors, so the pre-run function is used
			cmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
				return err
			}
			return
		}
		os.Exit(0)
	})
	return nil
}

// GenCompletion is a helper function that generates a shell completion
// script for the provided shell.
func GenCompletion(cmd *cobra.Command, shell Shell) error {
	switch shell {
	case Bash:
		return cmd.Root().GenBashCompletion(cmd.OutOrStdout())
	case Zsh:
		return cmd.Root().GenZshCompletion(cmd.OutOrStdout())
	case Fish:
		return cmd.Root().GenFishCompletion(cmd.OutOrStdout(), true)
	case PowerShell:
		return cmd.Root().GenPowerShellCompletionWithDesc(cmd.OutOrStdout())
	default:
		return fmt.Errorf("%w: %s", ErrInvalidShell, shell)
	}
}
