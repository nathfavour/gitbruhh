# gitbruhh

`gitbruhh` is a robust, modular, and intelligent CLI tool for fetching information from GitHub without leaving your terminal.

## Features

- **User Info**: Fetch detailed profile information about any GitHub user.
- **Repository Info**: Get stats, description, and status of any public repository.
- **Search**: Search for repositories using keywords.
- **Issues**: List recent open issues for any repository.
- **Authentication**: Supports Personal Access Tokens (PAT) for higher rate limits and private data access.

## Installation

```bash
# Clone the repository
git clone https://github.com/nathfavour/gitbruhh
cd gitbruhh

# Build the tool
go build -o gitbruhh main.go

# Move to your PATH (optional)
mv gitbruhh /usr/local/bin/
```

## Usage

### Get User Info
```bash
gitbruhh user [username]
```

### Get Repository Info
```bash
gitbruhh repo [owner/repo]
```

### Search Repositories
```bash
gitbruhh search [query]
```

### List Recent Issues
```bash
gitbruhh issues [owner/repo]
```

## Authentication

By default, `gitbruhh` works unauthenticated (subject to GitHub's rate limits for public APIs). To increase your rate limit or access private information, set the `GITHUB_TOKEN` environment variable:

```bash
export GITHUB_TOKEN=your_personal_access_token
gitbruhh user nathfavour
```

## Architecture

- `cmd/`: CLI command definitions using Cobra.
- `pkg/gh/`: Core logic for interacting with the GitHub API.
- `pkg/ui/`: Formatting and display logic.
