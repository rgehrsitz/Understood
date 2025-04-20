package cache

import (
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
)

// GetSHA1 returns the SHA-1 hash of a file's contents.
func GetSHA1(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	sum := sha1.Sum(data)
	return hex.EncodeToString(sum[:]), nil
}

// Simple file-based cache: summaries stored as .cache/<sha1>.txt
func GetFromCache(cacheDir, sha string) (string, error) {
	path := filepath.Join(cacheDir, sha+".txt")
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func PutToCache(cacheDir, sha, summary string) error {
	os.MkdirAll(cacheDir, 0o755)
	path := filepath.Join(cacheDir, sha+".txt")
	return ioutil.WriteFile(path, []byte(summary), 0o644)
}
