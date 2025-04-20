package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var repo, out, format, model string

	rootCmd := &cobra.Command{
		Use:   "understood",
		Short: "Generate LLM-powered documentation for public repos",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("[Understood] CLI stub. Flags:")
			fmt.Println("repo:", repo)
			fmt.Println("out:", out)
			fmt.Println("format:", format)
			fmt.Println("model:", model)
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
