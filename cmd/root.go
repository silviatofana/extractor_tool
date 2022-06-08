package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

type rootConfig struct {
	SkipLibraries *bool
	SkipUpdate    *bool
	Seeds         *[]string
	Emails        *[]string
	GitPath       *string
	OutPutPath    *string
	HashImportant *bool
}

var (
	rootCmd = &cobra.Command{
		Use:   "extractor_tool",
		Short: "Extract data from a Git repository",
		Long: `Use this command to extract and upload repo data your CodersRank profile.
Example usage: extractor_tool path --repo_path /path/to/repo`,
	}

	RootConfig rootConfig

	emailString *string
	seedsString *string
	Version     string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootConfig.SkipLibraries = rootCmd.PersistentFlags().Bool("skip_libraries", false, "Turns off the library detection in order to reduce the execution time")
	RootConfig.SkipUpdate = rootCmd.PersistentFlags().Bool("skip_update", false, "If set the auto-update is skipped")
	emailString = rootCmd.PersistentFlags().String("emails", "", "Predefined emails. Example: \"alim.giray@codersrank.io,alimgiray@gmail.com\"")
	seedsString = rootCmd.PersistentFlags().String("seeds", "", "The seed is used to find similar emails. Example: \"alimgiray, alimgiray@codersrank.io\"")
	RootConfig.GitPath = rootCmd.PersistentFlags().String("git_path", "", "where the Git binary is")
	RootConfig.OutPutPath = rootCmd.PersistentFlags().String("output_path", "./export", "Where to put output file. Existing exports will be overwritten.")
	RootConfig.HashImportant = rootCmd.PersistentFlags().Bool("hash_important", false, "Emails will be hashed.")
}

func initConfig() {
	emails := make([]string, 0)
	if len(*emailString) > 0 {
		emails = strings.Split(*emailString, ",")
	}
	RootConfig.Emails = &emails

	seeds := make([]string, 0)
	if len(*seedsString) > 0 {
		seeds = strings.Split(*seedsString, ",")
	}

	RootConfig.Seeds = &seeds

	// Find git executable if it is not provided
	if *RootConfig.GitPath == "" {
		gitPath, err := exec.LookPath("git")
		if err != nil {
			defaultGitPath := "/usr/bin/git"
			fmt.Printf("Couldn't find git path. Fall back to default (%s). Error: %s.\n", defaultGitPath, err.Error())
			// Try default git path
			*RootConfig.GitPath = defaultGitPath
			return
		}
		gitPath = strings.TrimRight(gitPath, "\r\n")
		gitPath = strings.TrimRight(gitPath, "\n")

		*RootConfig.GitPath = gitPath
	}
}
