package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/joho/godotenv"
)

//go:embed docs/*.md
var embeddedDocs embed.FS

//go:embed templates/theme.html
var themeTemplate string

//go:embed static/css/* static/js/*
var staticFiles embed.FS

const defaultPort = "3005"

type PageData struct {
	Title   string
	Content template.HTML
	Pages   []PageLink
}

type PageLink struct {
	Name string
	Path string
}

// titleCase converts a string to title case (first letter of each word capitalized)
func titleCase(s string) string {
	if s == "" {
		return s
	}
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

func main() {
	// Load .env file if it exists
	godotenv.Load()

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Write embedded static files to dist
	if err := writeStaticFiles(); err != nil {
		log.Fatalf("Error writing static files: %v", err)
	}

	// Convert markdown files to HTML
	if err := convertMarkdownFiles(); err != nil {
		log.Fatalf("Error converting markdown files: %v", err)
	}

	// Setup HTTP server
	http.HandleFunc("/", serveHTML)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("dist/static"))))

	log.Printf("Starting server on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func writeStaticFiles() error {
	// Create dist/static directories
	if err := os.MkdirAll("dist/static", 0755); err != nil {
		return fmt.Errorf("failed to create dist/static directory: %w", err)
	}

	// Write CSS files
	cssFiles := []string{"pico.min.css"}
	for _, file := range cssFiles {
		content, err := staticFiles.ReadFile("static/css/" + file)
		if err != nil {
			return fmt.Errorf("failed to read embedded CSS %s: %w", file, err)
		}
		cssPath := filepath.Join("dist", "static", file)
		if err := os.WriteFile(cssPath, content, 0644); err != nil {
			return fmt.Errorf("failed to write CSS %s: %w", file, err)
		}
		log.Printf("Wrote static file: %s", file)
	}

	// Write JS files
	jsFiles := []string{"vue.global.prod.js"}
	for _, file := range jsFiles {
		content, err := staticFiles.ReadFile("static/js/" + file)
		if err != nil {
			return fmt.Errorf("failed to read embedded JS %s: %w", file, err)
		}
		jsPath := filepath.Join("dist", "static", file)
		if err := os.WriteFile(jsPath, content, 0644); err != nil {
			return fmt.Errorf("failed to write JS %s: %w", file, err)
		}
		log.Printf("Wrote static file: %s", file)
	}

	return nil
}

func convertMarkdownFiles() error {
	// Create dist directory
	if err := os.MkdirAll("dist", 0755); err != nil {
		return fmt.Errorf("failed to create dist directory: %w", err)
	}

	// Try to read from local docs directory first
	localDocs := "docs"
	var mdFiles []string
	var docsSource string

	if _, err := os.Stat(localDocs); err == nil {
		// Local docs directory exists
		docsSource = "local"
		entries, err := os.ReadDir(localDocs)
		if err != nil {
			return fmt.Errorf("failed to read local docs directory: %w", err)
		}
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
				mdFiles = append(mdFiles, entry.Name())
			}
		}
	} else {
		// Use embedded docs
		docsSource = "embedded"
		entries, err := embeddedDocs.ReadDir("docs")
		if err != nil {
			return fmt.Errorf("failed to read embedded docs directory: %w", err)
		}
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
				mdFiles = append(mdFiles, entry.Name())
			}
		}
	}

	if len(mdFiles) == 0 {
		log.Println("Warning: No markdown files found")
		return nil
	}

	log.Printf("Converting %d markdown files from %s docs...", len(mdFiles), docsSource)

	// Build page links
	var pages []PageLink
	for _, filename := range mdFiles {
		name := strings.TrimSuffix(filename, ".md")
		name = strings.ReplaceAll(name, "-", " ")
		name = titleCase(name)
		htmlFilename := strings.TrimSuffix(filename, ".md") + ".html"
		pages = append(pages, PageLink{
			Name: name,
			Path: "/" + htmlFilename,
		})
	}

	// Convert each markdown file
	for _, filename := range mdFiles {
		var content []byte
		var err error

		if docsSource == "local" {
			content, err = os.ReadFile(filepath.Join(localDocs, filename))
		} else {
			content, err = embeddedDocs.ReadFile(filepath.Join("docs", filename))
		}

		if err != nil {
			return fmt.Errorf("failed to read %s: %w", filename, err)
		}

		// Convert markdown to HTML
		htmlContent := markdownToHTML(content)

		// Get title from filename
		title := strings.TrimSuffix(filename, ".md")
		title = strings.ReplaceAll(title, "-", " ")
		title = titleCase(title)

		// Write HTML file
		htmlFilename := strings.TrimSuffix(filename, ".md") + ".html"
		htmlPath := filepath.Join("dist", htmlFilename)

		// Create full page HTML with sidebar
		pageData := PageData{
			Title:   title,
			Content: template.HTML(htmlContent),
			Pages:   pages,
		}

		tmpl, err := template.New("page").Parse(themeTemplate)
		if err != nil {
			return fmt.Errorf("failed to parse template: %w", err)
		}

		file, err := os.Create(htmlPath)
		if err != nil {
			return fmt.Errorf("failed to create %s: %w", htmlPath, err)
		}
		defer file.Close()

		if err := tmpl.Execute(file, pageData); err != nil {
			return fmt.Errorf("failed to execute template for %s: %w", htmlFilename, err)
		}

		log.Printf("Converted %s -> %s", filename, htmlFilename)
	}

	return nil
}

func markdownToHTML(md []byte) string {
	// Create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// Create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer))
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		path = "/index.html"
	}

	// Remove leading slash
	path = strings.TrimPrefix(path, "/")

	// Check if file exists in dist directory
	filePath := filepath.Join("dist", path)
	if _, err := os.Stat(filePath); err == nil {
		// Serve the HTML file
		content, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, "Failed to read file", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(content)
		return
	}

	// File not found, show directory listing
	entries, err := os.ReadDir("dist")
	if err != nil {
		http.Error(w, "Failed to read directory", http.StatusInternalServerError)
		return
	}

	var pages []PageLink
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".html") {
			name := strings.TrimSuffix(entry.Name(), ".html")
			name = strings.ReplaceAll(name, "-", " ")
			name = titleCase(name)
			pages = append(pages, PageLink{
				Name: name,
				Path: "/" + entry.Name(),
			})
		}
	}

	// Render index page with list of pages
	pageData := PageData{
		Title:   "Documentation",
		Content: template.HTML("<p>Available documentation pages:</p>"),
		Pages:   pages,
	}

	tmpl, err := template.New("page").Parse(themeTemplate)
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, pageData); err != nil {
		log.Printf("Error executing template: %v", err)
	}
}
