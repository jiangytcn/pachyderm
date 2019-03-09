package cmdutil

import (
	"fmt"
	"os"
	"strings"

	"github.com/pachyderm/pachyderm/src/client/pfs"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// RunFixedArgs wraps a function in a function
// that checks its exact argument count.
func RunFixedArgs(numArgs int, run func([]string) error) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		if len(args) != numArgs {
			fmt.Printf("expected %d arguments, got %d\n\n", numArgs, len(args))
			cmd.Usage()
		} else {
			if err := run(args); err != nil {
				ErrorAndExit("%v", err)
			}
		}
	}
}

// RunBoundedArgs wraps a function in a function
// that checks its argument count is within a range.
func RunBoundedArgs(min int, max int, run func([]string) error) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		if len(args) < min || len(args) > max {
			fmt.Printf("expected %d to %d arguments, got %d\n\n", min, max, len(args))
			cmd.Usage()
		} else {
			if err := run(args); err != nil {
				ErrorAndExit("%v", err)
			}
		}
	}
}

// Run makes a new cobra run function that wraps the given function.
func Run(run func(args []string) error) func(*cobra.Command, []string) {
	return func(_ *cobra.Command, args []string) {
		if err := run(args); err != nil {
			ErrorAndExit(err.Error())
		}
	}
}

// ErrorAndExit errors with the given format and args, and then exits.
func ErrorAndExit(format string, args ...interface{}) {
	if errString := strings.TrimSpace(fmt.Sprintf(format, args...)); errString != "" {
		fmt.Fprintf(os.Stderr, "%s\n", errString)
	}
	os.Exit(1)
}

// ParseCommits takes a slice of arguments of the form "repo/commit-id" or
// "repo" (in which case we consider the commit ID to be empty), and returns
// a list of *pfs.Commits
func ParseCommits(args []string) ([]*pfs.Commit, error) {
	var commits []*pfs.Commit
	for _, arg := range args {
		parts := strings.SplitN(arg, "/", 2)
		hasRepo := len(parts) > 0 && parts[0] != ""
		hasCommit := len(parts) == 2 && parts[1] != ""
		if hasCommit && !hasRepo {
			return nil, fmt.Errorf("invalid commit id \"%s\": repo cannot be empty", arg)
		}
		commit := &pfs.Commit{
			Repo: &pfs.Repo{
				Name: parts[0],
			},
		}
		if len(parts) == 2 {
			commit.ID = parts[1]
		} else {
			commit.ID = ""
		}
		commits = append(commits, commit)
	}
	return commits, nil
}

// ParseBranches takes a slice of arguments of the form "repo/branch-name" or
// "repo" (in which case we consider the branch name to be empty), and returns
// a list of *pfs.Branches
func ParseBranches(args []string) ([]*pfs.Branch, error) {
	commits, err := ParseCommits(args)
	if err != nil {
		return nil, err
	}
	var result []*pfs.Branch
	for _, commit := range commits {
		result = append(result, &pfs.Branch{Repo: commit.Repo, Name: commit.ID})
	}
	return result, nil
}

// RepeatedStringArg is an alias for []string
type RepeatedStringArg []string

func (r *RepeatedStringArg) String() string {
	result := "["
	for i, s := range *r {
		if i != 0 {
			result += ", "
		}
		result += s
	}
	return result + "]"
}

// Set adds a string to r
func (r *RepeatedStringArg) Set(s string) error {
	*r = append(*r, s)
	return nil
}

// Type returns the string representation of the type of r
func (r *RepeatedStringArg) Type() string {
	return "[]string"
}

type BatchArgs struct {
	Positionals []string
	Flags map[string][]string
}

func (args *BatchArgs) GetStringFlag(name string) string {
	if len(args.Flags[name]) > 0 {
		return args.Flags[name][0]
	}
	return ""
}

func partitionBatchArgs(args []string, positionalCount int) [][]string {
	// Partition by sets of positional args, flags apply to the previous positional args
	sets := [][]string{{}}
	index := 0
	count := 0

	for _, x := range args {
		if strings.HasPrefix(x, "-") {
			sets[index] = append(sets[index], x)
			// TODO: this only works because all existing batch flags require a
			// value, otherwise we'll need to check individual flags
			if !strings.Contains(x, "=") {
				count -= 1
			}
		} else if count < positionalCount {
			count += 1
			sets[index] = append(sets[index], x)
		} else {
			sets = append(sets, []string{x})
			index += 1
			count = 0
		}
	}

	return sets
}

// newBatchParserCommand creates a dummy command parser that inherits flags from
// the originalCommand, but is not attached to the main command tree.  This
// gives us a bit more freedom to iteratively parse a batch of commands without
// jumping through as many hoops.
func newBatchParserCommand(originalCommand *cobra.Command) *cobra.Command {
	// Create an inner command for parsing individual commands
	result := &cobra.Command{}
	result.Flags().AddFlagSet(originalCommand.LocalFlags())

	// Copy over inherited flags
	ancestor := originalCommand.Parent()
	for ancestor != nil {
		result.Flags().AddFlagSet(ancestor.Flags())
		ancestor = ancestor.Parent()
	}

	return result
}

// RunBatchCommand generates a run function for a batch CLI command.  This works
// the same as a normal command except that we parse any number of sets of
// arguments from the command-line args, each set corresponding to a request.
// These parsed sets will be passed to the specified 'run' callback, which
// can then be validated and handled.
func RunBatchCommand(
	positionalCount int,
	run func([]BatchArgs) error,
) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, originalArgs []string) {
		// As long as there's any test coverage, hopefully this will save some time
		if !cmd.DisableFlagParsing {
			panic("Batch commands must disable flag parsing")
		}

		// Partition the args based on the number of positional arguments
		args := partitionBatchArgs(originalArgs, positionalCount)

		// Set a run function that appends to our set of parsed args
		parsedArgs := []BatchArgs{}
		innerCommand := newBatchParserCommand(cmd)
		innerCommand.Run = func (runCmd *cobra.Command, runArgs []string) {
			parsedArgs = append(parsedArgs, BatchArgs{runArgs, map[string][]string{}})
		}

		for _, argSet := range args {
			flags := map[string][]string{}
			innerCommand.SetArgs(argSet)
			innerCommand.Execute()
			innerCommand.Flags().ParseAll(
				argSet,
				func (flag *pflag.Flag, value string) error {
					flags[flag.Name] = append(flags[flag.Name], value)
					return nil
				},
			)
			parsedArgs[len(parsedArgs) - 1].Flags = flags
		}

		if err := run(parsedArgs); err != nil {
			ErrorAndExit("%v", err)
		}
	}
}

// SetDocsUsage sets the usage string for a docs-style command.  Docs commands
// have no functionality except to output some docs and related commands, and
// should not specify a 'Run' attribute.
func SetDocsUsage(command *cobra.Command, subcommands []*cobra.Command) {
    command.SetUsageTemplate(`Usage:
  pachctl [command]{{if gt .Aliases 0}}

Aliases:
  {{.NameAndAliases}}
{{end}}{{if .HasExample}}

Examples:
{{ .Example }}{{end}}{{ if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimRightSpace}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsHelpCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}
`)

    command.SetHelpTemplate(`{{or .Long .Short}}
{{.UsageString}}`)

    // This song-and-dance is so that we can render the related commands without
    // actually having them usable as subcommands of the docs command.
    // That is, we don't want `pachctl job list-job` to work, it should just
    // be `pachctl list-job`.  Therefore, we lazily add/remove the subcommands
    // only when we try to render usage for the docs command.
    originalUsage := command.UsageFunc()
    command.SetUsageFunc(func (c *cobra.Command) error {
        newUsage := command.UsageFunc()
        command.SetUsageFunc(originalUsage)
        defer command.SetUsageFunc(newUsage)

        command.AddCommand(subcommands...)
        defer command.RemoveCommand(subcommands...)

        command.Usage()
        return nil
    })
}
