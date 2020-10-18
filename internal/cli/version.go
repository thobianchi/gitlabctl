package cli

import (
	"flag"
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/spf13/cobra"
)

// Version is dynamically set by the ci or overridden by the Makefile.
var Version = "DEV"

// BuildDate is dynamically set at build time by the cli or overridden in the Makefile.
var BuildDate = "" // YYYY-MM-DD

func init() {
	if Version == "DEV" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		}
	}
}

func VersionSubcommand(cmd *cobra.Command) {
	versionOutput := versionFormat(Version, BuildDate)
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show gitlabctl version",
		Long:  "Show gitlabctl version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(versionOutput)
		},
	}
	cmd.AddCommand(versionCmd)
}

// versionFormat return the version string nicely formatted
func versionFormat(version, buildDate string) string {
	if Version == "DEV" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		}
	}

	if buildDate != "" {
		version = fmt.Sprintf("%s (%s)", version, buildDate)
	}

	version = fmt.Sprintf("gitlabctl version: %s", version)
	// don't return GoVersion during a test run for consistent test output
	if flag.Lookup("test.v") != nil {
		return version
	}

	return fmt.Sprintf("%s, Go Version: %s", version, runtime.Version())
}
