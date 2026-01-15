# Getting Started

This guide will help you get started with the documentation system.

## Installation

### Prerequisites

- Go 1.21 or higher
- A text editor

### Steps

1. **Clone the repository**
   ```bash
   git clone https://github.com/scriptmaster/docs.git
   cd docs
   ```

2. **Build the project**
   ```bash
   go build -o docs.exe
   ```

3. **Run the server**
   ```bash
   ./docs.exe
   ```

## Configuration

You can configure the server port using environment variables:

### Using .env file

Create a `.env` file in the root directory:

```
PORT=8080
```

### Using environment variable

```bash
PORT=8080 ./docs.exe
```

## Adding Content

Simply add your markdown files to the `docs/` directory and they will be automatically converted to HTML.

> **Tip**: The server will use embedded docs if no local `docs/` directory is found.
