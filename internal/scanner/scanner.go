package scanner

import (
	"io/fs"
	"path/filepath"
	"strings"
)

var defaultExcludes = []string{".git", "node_modules", "vendor", "dist", "build", ".DS_Store"}

// Scan walks the repo and returns code/doc files, excluding noise.
func Scan(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		for _, ex := range defaultExcludes {
			if d.IsDir() && d.Name() == ex {
				return filepath.SkipDir
			}
		}
		if !d.IsDir() && isInterestingFile(d.Name()) {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func isInterestingFile(name string) bool {
	l := strings.ToLower(name)
	return strings.HasSuffix(l, ".go") || strings.HasSuffix(l, ".py") || strings.HasSuffix(l, ".js") || strings.HasSuffix(l, ".ts") || strings.HasSuffix(l, ".md") || strings.HasSuffix(l, ".rs") || strings.HasSuffix(l, ".java") || strings.HasSuffix(l, ".c") || strings.HasSuffix(l, ".cpp") || strings.HasSuffix(l, ".cs")
}
