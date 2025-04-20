package renderer

import (
	"html/template"
	"os"
)

type RepoSummary struct {
	RepoURL   string
	Summaries map[string]string // file -> summary
}

// RenderMarkdown renders summaries to Markdown using a template.
func RenderMarkdown(outPath string, repoURL string, summaries map[string]string) error {
	tmpl, err := template.ParseFiles("templates/repo.md.tmpl")
	if err != nil {
		return err
	}
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()
	return tmpl.Execute(f, RepoSummary{RepoURL: repoURL, Summaries: summaries})
}

// RenderHTML renders summaries to HTML using a template.
func RenderHTML(outPath string, repoURL string, summaries map[string]string) error {
	tmpl, err := template.ParseFiles("templates/repo.html.tmpl")
	if err != nil {
		return err
	}
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()
	return tmpl.Execute(f, RepoSummary{RepoURL: repoURL, Summaries: summaries})
}
