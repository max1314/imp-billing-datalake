// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

/*
Package sharedflag holds the FlagSet that is shared by all dynamic flags of Improbable Services.

Basic usage:

     myFlag = sharedflag.Set.Int("some_flag_name", "my_default", "dynamic flag for something")

When creating a top-level command that uses these flags, you must add them to the flag list
by calling AddToCommand.
*/
package sharedflags

import (
	"github.com/mwitkow/go-flagz/monitoring"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	// Set is the repository for all dynamic flags across all libraries of Improbable.
	Set = pflag.NewFlagSet("shared_flagz", pflag.ContinueOnError)

	// Flag to show complete list of shared flags in help.
	fFullHelp = Set.Bool("full_help", false, "show shared flags")
)

func init() {
	Set.Lookup("full_help").Hidden = true
}

func init() {
	// Register the monitoring of checksums for this shared flagset.
	monitoring.MustRegisterFlagSet("shared", Set)
}

func addFullHelpFlag(c *cobra.Command, addFlagDoc bool) {
	root := c.Root()

	previous := root.HelpFunc()
	root.SetHelpFunc(func(c *cobra.Command, args []string) {
		if *fFullHelp {
			Set.VisitAll(func(f *pflag.Flag) { f.Hidden = false })
		}
		previous(c, args)
	})

	if !addFlagDoc {
		return
	}

	root.SetUsageTemplate(root.UsageTemplate() + `Use "{{.CommandPath}} [command] --help --full_help" for a complete list of flags.
`)
}

// AddToCommand adds shared flags to the PersistentFlagSet of a command and hides them.
//
// This should NOT be used by servers.
//
// All sharedflags will be set to hidden when this is called.  All hidden shared flags will be shown if
// command help is invoked with --full_help, e.g.:
//   --help --full_help
//   help [command] --full_help
// etc.
//
// How to show specific flags for a tool or command
//
// Commands may want to un-hide individual flags that are of specific interest to the command. E.g.:
//   sharedflags.Set.Lookup("a_flag_i_want").Hidden = false
//
// If a sub-command has its own set of flags that aren't really shared, it should register them
// directly on that sub-command. E.g.:
//   myCmd.Flags().Bool("enable_thing", false, "enable doing that thing")
//
// If a tool has flags specific to the tool, but shared between commands, it may add them to a
// tool-specific FlagSet. E.g.:
//   var flags = pflag.NewFlagSet("tool_flags", pflag.ContinueOnError)
//   var enableThing = flags.Bool("enable_thing", false, "enable doing that thing")
//   func main() {
//     rootCmd := &cobra.Command{...}
//     rootCmd.PersistentFlagSet().AddFlagSet(flags)
//   }
//
func AddToCommand(c *cobra.Command) {
	Set.VisitAll(func(f *pflag.Flag) { f.Hidden = true })
	AddToCommandWithFullHelp(c)
	addFullHelpFlag(c, true)
}

// AddToCommandWithHidden adds shared flags to the PersistenFlagSet of a command and hides them.
// This version does not document the --full_help flag, and is intended for tools shipped to external users.
func AddToCommandWithHidden(c *cobra.Command) {
	Set.VisitAll(func(f *pflag.Flag) { f.Hidden = true })
	AddToCommandWithFullHelp(c)
	addFullHelpFlag(c, false)
}

// AddToCommandWithFullFlags adds flags to the PersistentFlagSet of the command c.
func AddToCommandWithFullHelp(c *cobra.Command) {
	c.Root().PersistentFlags().AddFlagSet(Set)
	previous := c.Root().PersistentPreRunE
	c.Root().PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// Run sharedflag.Set's Parse so appropriate properties are set on sharedflags. This should be allowed to fail,
		// since at this point the `args` are already parsed and may contain the positional arguments that look like
		// flags (e.g. start with a dash), but are allowed here since they were passed after a `--` end-of-options
		// delimiter which explicitly signals the end of options and disables further option processing. For more details
		// check https://linux.die.net/man/1/bash and pflag source.
		_ = Set.Parse(args)
		if previous != nil {
			return previous(cmd, args)
		}
		return nil
	}
}
