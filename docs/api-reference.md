# API Reference

Complete API documentation for the project.

## Core Functions

### convertMarkdownFiles()

Converts all markdown files in the `docs/` directory to HTML files in the `dist/` directory.

**Returns**: `error`

### markdownToHTML(md []byte)

Converts markdown content to HTML.

**Parameters**:
- `md []byte` - The markdown content as bytes

**Returns**: `string` - The HTML output

### serveHTML(w http.ResponseWriter, r *http.Request)

HTTP handler for serving HTML files.

**Parameters**:
- `w http.ResponseWriter` - The response writer
- `r *http.Request` - The HTTP request

## Data Types

### PageData

```go
type PageData struct {
    Title   string
    Content template.HTML
    Pages   []PageLink
}
```

### PageLink

```go
type PageLink struct {
    Name string
    Path string
}
```

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| PORT     | 3005    | Server port |

## Examples

### Basic Usage

```go
// Start the server
go run main.go
```

### Custom Port

```bash
PORT=8080 go run main.go
```
