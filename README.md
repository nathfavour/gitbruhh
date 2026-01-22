# xoy

`xoy` is a robust, modular, and intelligent CLI tool for fetching information from GitHub without leaving your terminal.

## Features

- **User Info**: Fetch detailed profile information about any GitHub user.
- **Repository Info**: Get stats, description, and status of any public repository.
- **Search**: Search for repositories using keywords.
- **Issues**: List recent open issues for any repository.
- **Multi-Provider Architecture**:
  - **API**: Uses the official GitHub REST API.
  - **Scraper**: Falls back to web scraping if the API is unavailable or rate-limited.
  - **Failsafe**: Automatically switches between providers to ensure you always get the data.
- **Authentication**: Optional. Works completely without login by falling back to scraping or using unauthenticated API calls.

## Installation

```bash
# Clone the repository
git clone https://github.com/nathfavour/xoy
cd xoy

# Build the tool
go build -o xoy main.go
```

## Usage

### Get User Info
```bash
./xoy user [username]
```

### Get Repository Info
```bash
./xoy repo [owner/repo]
```

### Search Repositories
```bash
./xoy search [query]
```

### List Recent Issues
```bash
./xoy issues [owner/repo]
```

### Provider Selection
You can explicitly choose a provider or let `xoy` decide automatically:
```bash
./xoy repo google/go-github --provider scraper
./xoy repo google/go-github --provider api
./xoy repo google/go-github --provider auto # default
```

## Authentication

By default, `xoy` works unauthenticated. To increase your rate limit for the API provider, set the `GITHUB_TOKEN` environment variable:

```bash
export GITHUB_TOKEN=your_personal_access_token
./xoy user nathfavour
```

## Architecture

- `cmd/`: CLI command definitions using Cobra.
- `pkg/gh/`: Core logic for interacting with GitHub (API and Scraping).
- `pkg/ui/`: Formatting and display logic.