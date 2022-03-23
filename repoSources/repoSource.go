package repoSource

import (
	"fmt"
	"io/ioutil"

	"github.com/Techloopio/extractor_tool/entities"
	"github.com/Techloopio/extractor_tool/extractor"
)

type ExtractConfig struct {
	OutputPath    string
	GitPath       string
	HashImportant bool
	UserEmails    []string
	Seeds         []string
	SkipLibraries bool
}

// RepoSource describes the interface that each provider has to implement
type RepoSource interface {
	// GetRepos provides the list of the repositories from the given provider
	GetRepos() []*entities.Repository
	// Clone clones the given repository to the given directory
	// returns with the cloned path and an error if any.
	Clone(repository *entities.Repository) (string, error)
	// CleanUp revert to the state before extraction.
	// E.g. remove temporary files, directories.
	CleanUp()
}

func ExtractFromSource(source RepoSource, config ExtractConfig) error {
	repos := source.GetRepos()

	if config.OutputPath == "" {
		outputDir, err := ioutil.TempDir("", "clone_dir_")
		if err != nil {
			return fmt.Errorf("couldn't create temp dir for artifacts. Try to set it with --output_path. Error: %s", err.Error())
		}
		config.OutputPath = outputDir
	}

	for _, repo := range repos {
		path, err := source.Clone(repo)
		if err != nil {
			fmt.Println("Couldn't clone repository. Error:", err.Error())
		}

		repoExtractor := extractor.RepoExtractor{
			RepoPath:      path,
			OutputPath:    config.OutputPath + "/" + repo.GetSafeFullName(),
			GitPath:       config.GitPath,
			HashImportant: config.HashImportant,
			UserEmails:    config.UserEmails,
			Seed:          config.Seeds,
			SkipLibraries: config.SkipLibraries,
		}

		err = repoExtractor.Extract()
		if err != nil {
			fmt.Println("Error during execution.", err.Error())
			continue
		}

	}
	source.CleanUp()

	return nil
}
