# Documentation Server

A lightweight, fast documentation server written in Go that converts Markdown files to beautiful HTML pages with a Material-inspired theme.

## Features

- 📝 **Markdown to HTML Conversion**: Automatically converts `docs/*.md` files to `dist/*.html`
- 🎨 **Material-Inspired Theme**: Beautiful design using Pico CSS
- ⚡ **Fast & Lightweight**: Built with Go for optimal performance
- 🔌 **Embedded Support**: Can use embedded docs or local `docs/` directory
- 🌐 **Built-in Web Server**: Serves documentation on port 3005 (configurable)
- 🎯 **Vue.js Integration**: Enhanced with Vue.js for interactive features
- 📱 **Responsive Design**: Works great on desktop and mobile

## Quick Start

### Download

Download the latest `docs.exe` for Windows from the releases.

### Run

1. Place `docs.exe` in your project directory
2. Create a `docs/` folder with your `.md` files
3. Run `docs.exe`
4. Open http://localhost:3005 in your browser

## Configuration

### Port Configuration

You can change the server port using:

**Environment Variable:**
```bash
PORT=8080 docs.exe
```

**Or create a `.env` file:**
```
PORT=8080
```

## Building from Source

### Prerequisites

- Go 1.21 or higher

### Build

```bash
# For Windows
GOOS=windows GOARCH=amd64 go build -o docs.exe

# For Linux
go build -o docserver

# For macOS
GOOS=darwin GOARCH=amd64 go build -o docserver
```

## Usage

### Directory Structure

```
your-project/
├── docs/
│   ├── index.md
│   ├── getting-started.md
│   └── api-reference.md
└── docs.exe
```

### Running the Server

```bash
# Windows
docs.exe

# Linux/macOS
./docserver
```

The server will:
1. Convert all `.md` files in `docs/` to HTML
2. Save them in the `dist/` directory
3. Start a web server on port 3005 (or configured port)
4. Serve the documentation with a beautiful Material-inspired theme

### Embedded Documentation

If no `docs/` directory exists, the server will use embedded documentation files (if built with embedded resources).

## Technology Stack

- **Backend**: Go
- **Markdown Parser**: gomarkdown/markdown
- **CSS Framework**: Pico CSS (via CDN)
- **JavaScript Framework**: Vue.js 3 (via CDN)
- **Theme**: Material-inspired design

## License

MIT

