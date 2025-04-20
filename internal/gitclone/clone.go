package gitclone

import (
	"fmt"
	"os"
	"strings"

	git "github.com/go-git/go-git/v5"
)

// CloneOrOpen clones a public repo if URL, or verifies a local path.
func CloneOrOpen(repo string, dest string) (string, error) {
	if strings.HasPrefix(repo, "http://") || strings.HasPrefix(repo, "https://") {
		fmt.Println("Cloning repo:", repo)
		_, err := git.PlainClone(dest, false, &git.CloneOptions{
			URL:      repo,
			Progress: os.Stdout,
		})
		if err != nil {
			return "", fmt.Errorf("clone failed: %w", err)
		}
		return dest, nil
	}
	// Local path
	if _, err := os.Stat(repo); os.IsNotExist(err) {
		return "", fmt.Errorf("local repo path does not exist: %s", repo)
	}
	return repo, nil
}
