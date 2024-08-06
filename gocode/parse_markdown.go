package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
)

func main() {
    // Slice to hold the article folders
    var articleFolders []string

    // Walk through the current directory to find article folders
    err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() && strings.HasPrefix(info.Name(), "article") {
            articleFolders = append(articleFolders, path)
        }
        return nil
    })
    if err != nil {
        log.Fatalf("Error walking the path: %v", err)
    }

    // Create a slice to hold all article HTML sections
    var articleHTMLSections []string

    for _, folder := range articleFolders {
        markdownFile := filepath.Join(folder, "index.md")
        htmlFile := filepath.Join(folder, "index.html")

        // Read markdown file
        input, err := os.ReadFile(markdownFile)
        if err != nil {
            log.Fatalf("Error reading markdown file: %v", err)
        }

        // Extract title, summary, and full content
        title, summary, content := extractContent(input)

        // Generate the article HTML section
        articleHTML := fmt.Sprintf(`
<div class="article">
    <h2>%s</h2>
    <div class="summary">%s</div>
    <div class="content">%s</div>
    <br>
    <span class="separator">More</span>
</div> <br>`, title, summary, content)

        // Append the article HTML section to the slice
        articleHTMLSections = append(articleHTMLSections, articleHTML)

        // Write the article HTML section to the article's index.html
        if err := ioutil.WriteFile(htmlFile, []byte(articleHTML), 0644); err != nil {
            log.Fatalf("Error writing HTML file: %v", err)
        }
    }

    // Insert all article HTML sections into the global index.html
    insertArticlesIntoIndex(articleHTMLSections)
}

func extractContent(markdownContent []byte) (string, string, string) {
    content := string(markdownContent)
    lines := strings.Split(content, "\n")

    var title, summary, restContent string
    var summaryHTML, restContentHTML []byte

    // Extract title
    if len(lines) > 0 && strings.HasPrefix(lines[0], "# ") {
        title = strings.TrimPrefix(lines[0], "# ")
        lines = lines[1:]
    }

    // Extract summary
    if len(lines) > 0 {
        summary = ""
        for _, line := range lines {
            if line != "" {
                summary = line
                break
            }
        }
        lines = lines[1:]
        summaryHTML = markdown.ToHTML([]byte(summary), nil, nil)
    }

    // Extract rest of the content
    if len(lines) > 0 {
        restContent = strings.Join(lines, "\n")
        restContentHTML = markdown.ToHTML([]byte(restContent), nil, nil)
    }

    return title, string(summaryHTML), string(restContentHTML)
}

func insertArticlesIntoIndex(articles []string) {
    // Read the global index.html template
    indexTemplate, err := ioutil.ReadFile("index_template.html")
    if err != nil {
        log.Fatalf("Error reading index template file: %v", err)
    }

    // Join all article HTML sections
    allArticlesHTML := strings.Join(articles, "\n")

    // Replace the placeholder with actual article sections
    finalIndexHTML := strings.Replace(string(indexTemplate), "{{articles}}", allArticlesHTML, 1)

    // Write the final index.html
    if err := ioutil.WriteFile("index.html", []byte(finalIndexHTML), 0644); err != nil {
        log.Fatalf("Error writing final index.html: %v", err)
    }
}

