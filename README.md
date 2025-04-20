# Understood

Generate beautiful, LLM-powered documentation for any public GitHub or local repo. Outputs Markdown or HTML overviews, per-file summaries, and more.

---

## Features
- **Public repo support:** Clone any public GitHub repo or use a local path.
- **LLM-powered summaries:** Uses OpenAI (gpt-4o, gpt-4, etc.) or Anthropic models to summarize code, docs, and structure.
- **Per-file and repo overviews:** Each file is summarized, and a repo-level overview is generated.
- **Markdown or HTML output:** Clean, customizable templates.
- **Caching:** Summaries are cached by file hash for speed and efficiency.
- **Noise filtering:** Ignores vendor, build, and other non-source directories by default.

## Quick Start

1. **Build the CLI:**
   ```sh
   go build ./cmd/understood
   ```

2. **Run on a public repo:**
   ```sh
   ./understood --repo https://github.com/gorilla/mux --out mux.md --format md --model gpt-4o
   ```
   - Requires an `OPENAI_API_KEY` in your environment.

3. **Output:**
   - Generates `mux.md` (or `.html`) with summaries of all interesting files in the repo.

## Usage

```
understood --repo <repo-url-or-path> --out <output-file> --format <md|html> --model <gpt-4o|claude-3> [flags]
```

- `--repo`   : Public GitHub URL or local repo path (required)
- `--out`    : Output file (default: `output.md`)
- `--format` : `md` or `html` (default: `md`)
- `--model`  : LLM model to use (default: `gpt-4o`)

## Example

```sh
./understood --repo https://github.com/gorilla/mux --out mux.md --format md --model gpt-4o
```

## Requirements
- Go 1.21+
- OpenAI API key (for summarization)

## Roadmap / TODO
- [ ] Retry/backoff for LLM rate limits
- [ ] Parallel summarization with throttling
- [ ] Progress bar/spinner
- [ ] Dry-run and verbose modes
- [ ] Custom template support
- [ ] Language-specific AST hooks
- [ ] Unit and integration tests

## License
MIT

---

PRs and suggestions welcome!
