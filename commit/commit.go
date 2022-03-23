package commit

type Commit struct {
	Hash         string
	AuthorName   string
	AuthorEmail  string
	Date         string
	ChangedFiles []*ChangedFile
	Libraries    map[string][]string
}

type OptimizedCommitForExport struct {
	AuthorEmail string              `json:"authorEmail"`
	Date        string              `json:"date"`
	Languages   []string            `json:"languages"`
	Insertions  int                 `json:"insertions"`
	Deletions   int                 `json:"deletions"`
	Libraries   map[string][]string `json:"libraries"`
	Commits     int                 `json:"commits"`
}

type ChangedFile struct {
	Path       string `json:"fileName"`
	Insertions int    `json:"insertions"`
	Deletions  int    `json:"deletions"`
	Language   string `json:"language"`
}
