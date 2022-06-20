## What is it?
Extractor is our new feature which captures all the commits you have done on certain projects
without accessing your code directly.
It creates a JSON file which gets the info from metadata stored in your commits.

Other information such as remote URLs, file names, emails, names is hashed. The Extractor will recognize if two commits belong to the same file but won’t know the files name. Moreover, the JSON output is saved on your computer and you can check for yourself what data is extracted and if it crosses your employers’ NDA before uploading it to EXP Timeline.
Once uploaded, the JSON file is classified as user content and is processed according to the [Terms and Conditions](https://www.lmc.eu/en/terms-of-services/specific-terms-of-services/techloop/).

## Dependencies
- [GO](https://go.dev/dl/)

## How to use it
The extractor_tool is written in Go, so you can either clone the repo and compile the program or just download the binary and start using it.
```
git clone --depth 1 https://github.com/Techloopio/extractor_tool.git
cd repo_info_extractor
go run . local --repo_path {relative path to repository}
```

Extracted JSON files are located in export folder (`./export/*`)

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
