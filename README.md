## What is it?
Extractor is our new feature which captures all the commits you have done on certain projects
without accessing your code directly.
It creates a JSON file which gets the info from metadata stored in your commits.

To avoid any fears about crossing the NDA this script extracts only the most important information from the repos:
- Number of inserted lines in each commit
- Number of deleted lines in each commit
Other information such as remote URLs, file names, emails, names are hashed. So we can know if two commits belong to the same file but we won't know the file name.
Moreover, the output is saved to your machine and you can check what data is extracted and you can decide whether you want to share it with us or not.


## Dependencies
- [GO](https://go.dev/dl/)

## How to use it
The extractor_tool is written in Go, so you can either clone the repo and compile the program or just download the binary and start using it.
```
git clone --depth 1 https://github.com/Techloopio/extractor_tool.git
cd repo_info_extractor
go run . local --repo_path ./path/to/repo
```

### Available commands
You can see the available commands and flags with the `--help` flag. For example:
```
./repo_info_extractor_osx --help
```
Commands:
-  `help` Help about any command
-  `local` Extract local repository by path
-  `version` Print the version number

The commands might have flags. For example `local` has:
`--repo-path` Path of the repo
