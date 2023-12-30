package cmd

import (
	"os"

	"github.com/devops-works/binenv/internal/app"
	"github.com/spf13/cobra"
)

// localCmd represents the local command
func downloadCmd(a *app.App) *cobra.Command {
	var targetOS, targetArch string

	cmd := &cobra.Command{
		Use:   "download [--lock] [--dry-run] [<distribution> <version> [<distribution> <version>]]",
		Short: "Download a version for the package",
		Long: `This command will download one or several distributions with the specified versions.
If --lock is used, versions from the .binenv.lock file in the current directory will be installed.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(1)
			}

			a.Download(targetOS, targetArch, args...)
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			switch len(args) % 2 {
			case 0:
				// complete application name
				return a.GetPackagesListWithPrefix(toComplete), cobra.ShellCompDirectiveNoFileComp
			case 1:
				// complete application version
				return a.GetAvailableVersionsFor(args[len(args)-1]), cobra.ShellCompDirectiveNoFileComp
			default:
				// huh ?
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
		},
	}

	cmd.Flags().StringVarP(&targetOS, "os", "o", "", "Target OS")
	cmd.Flags().StringVarP(&targetArch, "arch", "a", "", "Target architecture")

	return cmd
}
