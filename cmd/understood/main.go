package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/yourusername/understood/internal/gitclone"
	"github.com/yourusername/understood/internal/scanner"
	"github.com/yourusername/understood/internal/summarizer"
	"github.com/yourusername/understood/internal/cache"
	"github.com/yourusername/understood/internal/renderer"
)

func main() {
	var repo, out, format, model string

	rootCmd := &cobra.Command{
		Use:   "understood",
		Short: "Generate LLM-powered documentation for public repos",
		Run: func(cmd *cobra.Command, args []string) {
			if repo == "" {
				fmt.Fprintln(os.Stderr, "--repo is required")
				os.Exit(1)
			}
			fmt.Println("[Understood] Cloning or opening repo...")
			// 1. Clone or open repo
			tmpDir := "tmp_repo"
			repoPath, err := gitclone.CloneOrOpen(repo, tmpDir)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Repo error:", err)
				os.Exit(1)
			}

			fmt.Println("[Understood] Scanning files...")
			files, err := scanner.Scan(repoPath)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Scan error:", err)
				os.Exit(1)
			}
			if len(files) == 0 {
				fmt.Fprintln(os.Stderr, "No interesting files found.")
				os.Exit(1)
			}

			fmt.Printf("[Understood] Summarizing %d files...\n", len(files))
			summ := summarizer.NewSummarizer(model)
			summaries := make(map[string]string)
			cacheDir := ".cache"
			for i, f := range files {
				sha, _ := cache.GetSHA1(f)
				cached, err := cache.GetFromCache(cacheDir, sha)
				if err == nil && cached != "" {
					fmt.Printf("[%d/%d] (cache) %s\n", i+1, len(files), f)
					summaries[f] = cached
					continue
				}
				fmt.Printf("[%d/%d] Summarizing %s...\n", i+1, len(files), f)
				summary, err := summ.SummarizeFile(cmd.Context(), f)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Summarize error for %s: %v\n", f, err)
					summaries[f] = "[Summary failed]"
					continue
				}
				summaries[f] = summary
				cache.PutToCache(cacheDir, sha, summary)
			}

			fmt.Println("[Understood] Rendering output...")
			var renderErr error
			if format == "md" {
				renderErr = renderer.RenderMarkdown(out, repo, summaries)
			} else {
				renderErr = renderer.RenderHTML(out, repo, summaries)
			}
			if renderErr != nil {
				fmt.Fprintln(os.Stderr, "Render error:", renderErr)
				os.Exit(1)
			}
			fmt.Println("[Understood] Done! Output at:", out)
		},
	}

	rootCmd.Flags().StringVar(&repo, "repo", "", "Repo URL or local path")
	rootCmd.Flags().StringVar(&out, "out", "output.md", "Output file path")
	rootCmd.Flags().StringVar(&format, "format", "md", "Output format: md or html")
	rootCmd.Flags().StringVar(&model, "model", "gpt-4o", "LLM model to use")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
